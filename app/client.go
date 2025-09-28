package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	jquery "github.com/antchfx/jsonquery"
)

func setProxy(c http.Client, p ProxyConfig) *http.Client {
	// Attempt to parse the Proxy URL
	u, err := url.Parse(p.ProxyUrl)
	if err != nil {
		panic(err)
	}

	// Set the client transport to use the proxy
	c.Transport = &http.Transport{
		Proxy: http.ProxyURL(u),
	}

	fmt.Printf("Using proxy: %s\n", p.ProxyUrl)

	return &c
}

func makeClient(c ClientConfig) *http.Client {
	client := http.Client{
		Timeout: getDuration(c.Timeout),
	}

	if c.Proxy.Enable {
		fmt.Print("Setting proxy\n")
		setProxy(client, c.Proxy)
	}

	return &client
}

func request(u string) (resBody []byte, err error) {
	// First check the cache for the URL
	resBody, err = cache.Get([]byte(u))
	if err == nil {
		// If the URL is in the cache, return the cached response
		fmt.Printf("Found cached response for %s\n", u)
		return
	} else {
		fmt.Printf("Cache miss for %s\n", u)
		err = nil
	}

	// If the URL is not in the cache, make the request and cache the response
	fmt.Printf("Requesting %s\n", u)

	// Form the request
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return
	}

	// Make the request
	res, err := client.Do(req)
	if err != nil {
		return
	}
	fmt.Printf("Received response from %s: %s\n", u, res.Status)

	// Read the response body
	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}

	// Cache the response
	fmt.Printf("Caching response for %s, length: %d\n", u, len(resBody))
	err = cache.Set([]byte(u), resBody, cacheConfig.DefaultTTL)
	if err != nil {
		fmt.Printf("Error caching response: %s\n", err)
	}

	return
}

func parseResponse(b []byte) (res *jquery.Node, err error) {
	res, err = jquery.Parse(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("Error parsing response: %s\n", err)
	}
	return
}
