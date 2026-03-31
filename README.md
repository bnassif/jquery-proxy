# jQuery Proxy

A simple web proxy that allows users to perform jQuery operations against an HTTP(S) endpoint with caching.

## Overview

This application provides an OS-agnostic method for performing reliable jQuery expressions against remote HTTP(S) endpoints.

Originally designed to offer an easy method for parsing JSON data for firewalls to obtain IP lists for services published in JSON format, it enables dynamic loading of Alias contents.

### Details

The application accepts a URL and an xPath query to execute against JSON data retrieved from the URL. The parsed data is returned to the user, enabling JSON filtering without requiring a local package like `jq`.

Redis caching is supported and recommended to prevent too many outbound requests that may lead to rate-limiting.

## Deployment

The application is a single binary file that accepts the following flag arguments for runtime configuration:

| Flag                 | Default            | Description                                              |
|----------------------|--------------------|----------------------------------------------------------|
| `address`            | `0.0.0.0:8080`     | The address and port on which to listen.                 |
| `read-timeout`       | `5s`               | The server's read timeout.                               |
| `write-timeout`      | `10s`              | The server's write timeout.                              |
| `idle-timeout`       | `15s`              | The server's idle timeout.                               |
| `persistent-cookies` | `false`            | Determines whether to use persistent cookies.            |
| `timeout`            | `10s`              | The client's timeout duration.                           |
| `proxy-enable`       | `false`            | Specifies whether to use a forward proxy.                |
| `proxy-url`          |                    | The proxy URL to use (if `proxy-enable` is true).        |
| `cache-enable`       | `true`             | Specifies whether to use a cache.                        |
| `cache-ttl`          | `30`               | The default TTL (Time to Live) for the cache in seconds. |
| `cache-size`         | `50MB`             | The size of the cache in bytes.                          |

For details on deploying the application in Docker, see [Docker.md](DOCKER.md).