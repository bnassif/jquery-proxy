package app

import (
	"strconv"
	"time"

	units "github.com/docker/go-units"
)

type Config struct {
	Server ServerConfig
	Client ClientConfig
}

type ServerConfig struct {
	AddressPort  string
	ReadTimeout  string
	WriteTimeout string
	IdleTimeout  string
}

type ClientConfig struct {
	PersistentCookies bool
	Timeout           string
	Proxy             ProxyConfig
	Cache             CacheConfig
}

type CacheConfig struct {
	Enable     bool
	DefaultTTL int
	CacheSize  string
}

type ProxyConfig struct {
	Enable   bool
	ProxyUrl string
}

func getDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

func convertStorage(s string) (resp int64, err error) {
	// First try to return a raw integer
	resp, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		// Attempt to parse the storage size
		resp, err = units.FromHumanSize(s)
	}
	return
}
