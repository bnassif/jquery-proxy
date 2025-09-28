package main

import (
	"flag"

	"github.com/bnassif/jquery-proxy/app"
)

var config app.Config

func main() {
	// Parse the command line flags into a Config struct
	flag.StringVar(&config.Server.AddressPort, "address", "0.0.0.0:8080", "The address and port to listen on")
	flag.StringVar(&config.Server.ReadTimeout, "read-timeout", "5s", "The read timeout for the server")
	flag.StringVar(&config.Server.WriteTimeout, "write-timeout", "10s", "The write timeout for the server")
	flag.StringVar(&config.Server.IdleTimeout, "idle-timeout", "15s", "The idle timeout for the server")
	flag.BoolVar(&config.Client.PersistentCookies, "persistent-cookies", false, "Whether to use persistent cookies")
	flag.StringVar(&config.Client.Timeout, "timeout", "10s", "The client timeout")
	flag.BoolVar(&config.Client.Proxy.Enable, "proxy-enable", false, "Whether to use a proxy")
	flag.StringVar(&config.Client.Proxy.ProxyUrl, "proxy-url", "", "The proxy URL")
	flag.BoolVar(&config.Client.Cache.Enable, "cache-enable", true, "Whether to use a cache")
	flag.IntVar(&config.Client.Cache.DefaultTTL, "cache-ttl", 30, "The default TTL for the cache in seconds")
	flag.StringVar(&config.Client.Cache.CacheSize, "cache-size", "50MB", "The size of the cache; can be a string or an integer")
	showVersion := flag.Bool("v", false, "Print version and exit")

	flag.Parse()

	// If -v is supplied, print the version and exit
	if *showVersion {
		app.PrintVersion()
	}

	// Run the app with the specified configuration
	app.Run(config)
}
