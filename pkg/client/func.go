package client

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/bnassif/jquery-proxy/pkg/config"
	redis "github.com/redis/go-redis/v9"
)

func NewClient(config *config.ClientConfig, baseLogger *slog.Logger) *Client {
	c := Client{
		client: http.Client{
			Timeout: config.GetTimeout(),
		},
		config: *config,
		logger: baseLogger.With(slog.String("component", "client")),
	}

	if config.Proxy.Enable {
		c.setProxy(config.Proxy)
	}

	if config.Cache.Enable {
		rdb, err := newRedisClient(&config.Cache)
		if err != nil {
			panic(err)
		}
		c.redis = rdb
	}

	return &c
}

func newRedisClient(cfg *config.CacheConfig) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, err
	}

	if cfg.ConnectTimeout != "" {
		opt.DialTimeout = cfg.GetConnectTimeout()
	}
	if cfg.ReadTimeout != "" {
		opt.ReadTimeout = cfg.GetReadTimeout()
	}
	if cfg.WriteTimeout != "" {
		opt.WriteTimeout = cfg.GetWriteTimeout()
	}
	if cfg.PoolSize > 0 {
		opt.PoolSize = cfg.PoolSize
	}
	if cfg.MinIdleConns > 0 {
		opt.MinIdleConns = cfg.MinIdleConns
	}

	rdb := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
