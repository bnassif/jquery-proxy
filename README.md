# jQuery Proxy

A simple web proxy that allows end users to perform jQuery operations against an HTTP(S) endpoint with caching.

## Overview

This application provides an OS-agnostic method for performing reliable jQuery expressions against remote HTTP(S) endpoints.

Originally designed to offer an easy method for parsing and returning JSON data as standard HTML for XXsense firewall systems, it enables dynamic loading of Alias contents.

### Details

The application accepts a URL and an xPath query to execute against JSON data retrieved from the URL. The parsed data is returned to the user, enabling JSON filtering without requiring a local package like `jq`.

Before making a request, the application checks its built-in cache for a response, optimizing performance by returning cached responses when available.

## Usage

### Query Parameters

The parameters supplied to the application must be GET path parameters and must be URL-encoded.

| Name     | Required | Description                                                              |
|----------|----------|--------------------------------------------------------------------------|
| `url`    | Yes      | The target URL from which to pull JSON data for querying.                |
| `query`  | No       | The xPath query to execute against the JSON data.                        |
| `all`    | No       | Determines whether the application should perform `Query` or `QueryAll`. |

### Examples

#### GitHub API

- **url**: `https://api.github.com/meta`
- **query**: `hooks`

`http://<hostname>[:<port>]/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=<section e.g. hooks>`


#### Google API

- **url**: `https://www.gstatic.com/ipranges/goog.json`
- **query**: `//ipv4Prefix`
- **all**: `true`

`http://<hostname>[:<port>]/?url=https%3A%2F%2Fwww.gstatic.com%2Fipranges%2Fgoog.json&query=%2F%2Fipv4Prefix&all=true`

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