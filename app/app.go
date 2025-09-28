package app

import (
	"net/http"

	"github.com/coocood/freecache"
)

var client http.Client
var server http.Server
var cache freecache.Cache
var cacheConfig CacheConfig

func Run(config Config) {
	// Create a cache object, if caching is enabled
	if config.Client.Cache.Enable {
		cacheConfig = config.Client.Cache

		// Convert the cache size to an integer
		cacheSize, err := convertStorage(config.Client.Cache.CacheSize)
		if err != nil {
			panic(err)
		}
		cache = *freecache.NewCache(int(cacheSize))
	}

	server = *makeServer(config.Server)
	client = *makeClient(config.Client)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
