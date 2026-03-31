package cmdutil

import (
	"github.com/bnassif/jquery-proxy/pkg/config"
	"github.com/spf13/viper"
)

func BuildConfig(opts *viper.Viper) *config.Config {
	return &config.Config{
		Server: config.ServerConfig{
			Address:      opts.GetString("address"),
			Port:         opts.GetString("port"),
			ReadTimeout:  opts.GetString("read-timeout"),
			WriteTimeout: opts.GetString("write-timeout"),
			IdleTimeout:  opts.GetString("idle-timeout"),
		},
		Client: config.ClientConfig{
			PersistentCookies: opts.GetBool("persistent-cookies"),
			Timeout:           opts.GetString("client-timeout"),
			Cache: config.CacheConfig{
				Enable:         opts.GetBool("redis-enable"),
				URL:            opts.GetString("redis-url"),
				TTL:            opts.GetString("redis-ttl"),
				KeyPrefix:      opts.GetString("redis-prefix"),
				ConnectTimeout: opts.GetString("redis-ttl"),
				ReadTimeout:    opts.GetString("redis-connect-timeout"),
				WriteTimeout:   opts.GetString("redis-read-timeout"),
				PoolSize:       opts.GetInt("redis-write-timeout"),
				MinIdleConns:   opts.GetInt("redis-min-idle-connections"),
			},
			Proxy: config.ProxyConfig{
				Enable: opts.GetBool("proxy-enable"),
				URL:    opts.GetString("proxy-url"),
			},
		},
		Logging: config.LoggingConfig{
			Level:     opts.GetString("log-level"),
			Format:    opts.GetString("log-format"),
			AddSource: opts.GetBool("log-add-source"),
		},
	}
}
