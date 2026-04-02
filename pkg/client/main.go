package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	jquery "github.com/antchfx/jsonquery"
	"github.com/bnassif/jquery-proxy/pkg/config"

	redis "github.com/redis/go-redis/v9"
)

type Client struct {
	client http.Client
	redis  *redis.Client
	config config.ClientConfig
	logger *slog.Logger
}

func (c *Client) setProxy(proxyConfig config.ProxyConfig) {
	// Attempt to parse the Proxy URL
	u, err := url.Parse(proxyConfig.URL)
	if err != nil {
		panic(err)
	}

	// Set the client transport to use the proxy
	c.client.Transport = &http.Transport{
		Proxy: http.ProxyURL(u),
	}

	c.logger.Debug(
		"proxy enabled",
		slog.String("proxy_url", proxyConfig.URL),
	)
}

func (c *Client) cacheKey(u string) string {
	sum := sha256.Sum256([]byte(u))
	return c.config.Cache.KeyPrefix + hex.EncodeToString(sum[:])
}

func (c *Client) ReadCache(u string) []byte {
	if !c.config.Cache.Enable || c.redis == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	val, err := c.redis.Get(ctx, c.cacheKey(u)).Bytes()
	if err != nil {
		if err == redis.Nil {
			c.logger.Debug(
				"cache miss",
				slog.String("url", u),
			)
			return nil
		}
		c.logger.Error(
			"cache read fail",
			slog.String("url", u),
			slog.Any("error", err),
		)
		return nil
	}
	c.logger.Debug(
		"cache hit",
		slog.String("url", u),
	)
	return val
}

func (c *Client) WriteCache(u string, b []byte) {
	if !c.config.Cache.Enable || c.redis == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := c.redis.Set(ctx, c.cacheKey(u), b, c.config.Cache.GetTTL()).Err()
	if err != nil {
		c.logger.Error(
			"cache write fail",
			slog.String("url", u),
			slog.Any("error", err),
		)
	}
	return
}

func (c *Client) Parse(b []byte) (resBody *jquery.Node, err error) {
	resBody, err = jquery.Parse(bytes.NewReader(b))
	if err != nil {
		c.logger.Error(
			"json parsing fail",
			slog.Any("error", err),
		)
	}
	return
}

func (c *Client) Request(u string) (resBody []byte, err error) {
	resBody = c.ReadCache(u)

	if resBody == nil {
		req, err := http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			c.logger.Error(
				"request creation fail",
				slog.String("url", u),
				slog.Any("error", err),
			)
			return nil, err
		}

		rawResp, err := c.client.Do(req)
		if err != nil {
			c.logger.Error(
				"request fail",
				slog.String("url", u),
				slog.Any("error", err),
			)
			return nil, err
		}
		defer rawResp.Body.Close()

		c.logger.Debug(
			"response received",
			slog.String("url", u),
			slog.String("status", rawResp.Status),
			slog.Int("status_code", rawResp.StatusCode),
			slog.Int64("size", rawResp.ContentLength),
		)

		resBody, err = io.ReadAll(rawResp.Body)
		if err != nil {
			c.logger.Error(
				"body read failure",
				slog.String("url", u),
				slog.Any("error", err),
			)
			return nil, err
		}

		c.WriteCache(u, resBody)
	}
	return
}
