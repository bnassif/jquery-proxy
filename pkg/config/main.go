package config

import (
	"time"
)

type Config struct {
	Server  ServerConfig
	Client  ClientConfig
	Logging LoggingConfig
}

type ServerConfig struct {
	Address      string
	Port         string
	ReadTimeout  string
	WriteTimeout string
	IdleTimeout  string
}

func (c *ServerConfig) GetReadTimeout() time.Duration {
	return getDuration(c.ReadTimeout)
}

func (c *ServerConfig) GetWriteTimeout() time.Duration {
	return getDuration(c.WriteTimeout)
}

func (c *ServerConfig) GetIdleTimeout() time.Duration {
	return getDuration(c.IdleTimeout)
}

type ClientConfig struct {
	PersistentCookies bool
	Timeout           string
	Proxy             ProxyConfig
	Cache             CacheConfig
}

func (c *ClientConfig) GetTimeout() time.Duration {
	return getDuration(c.Timeout)
}

type CacheConfig struct {
	Enable    bool
	URL       string
	KeyPrefix string
	TTL       string

	ConnectTimeout string
	ReadTimeout    string
	WriteTimeout   string

	PoolSize     int
	MinIdleConns int
}

func (c *CacheConfig) GetTTL() time.Duration {
	return getDuration(c.TTL)
}

func (c *CacheConfig) GetConnectTimeout() time.Duration {
	return getDuration(c.ConnectTimeout)
}

func (c *CacheConfig) GetReadTimeout() time.Duration {
	return getDuration(c.ReadTimeout)
}

func (c *CacheConfig) GetWriteTimeout() time.Duration {
	return getDuration(c.WriteTimeout)
}

type ProxyConfig struct {
	Enable bool
	URL    string
}

type LoggingConfig struct {
	Level     string
	Format    string
	AddSource bool
}
