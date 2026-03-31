# Deployment

Deploying the jQuery Proxy app is made simple and quick.

## Docker

Dockerized deployments are quick and easy.

> [!TIP]
> It is recommended to deploy the container image to avoid unwanted environment parsing.

1. Copy the [docker-compose.yml](../deploy/docker-compose.yml) file contents to the host
2. Customize [the settings](#available-settings) as desired
3. Start the application with `docker compose up -d`

> [!NOTE]
> The provided compose file includes a redis service for configuration.

## Available Settings

The flags detailed below can also be specified as environment variables.

When specified as environment variables, they should be upper-case with hyphens replaced with underscores.  
For example, `--redis-read-timeout` would be set with `REDIS_READ_TIMEOUT`.

| Flag                           | Default   | Description                                                      |
|--------------------------------|-----------|------------------------------------------------------------------|
| `--address` / `-a`             | `0.0.0.0` | The address the server listens on                                |
| `--port` / `-p`                | `8080`    | The port the server listens on                                   |
| `--read-timeout`               | `5s`      | The read timeout for the server                                  |
| `--write-timeout`              | `10s`     | The write timeout for the server                                 |
| `--idle-timeout`               | `15s`     | The idle timeout for the server                                  |
| `--persistent-cookies`         | `false`   | Whether to use persistent cookies                                |
| `--client-timeout`             | `10s`     | The timeout for proxied requests                                 |
| `--redis-enable`               | `false`   | Whether to enable redis caching                                  |
| `--redis-url`                  |           | The url of the Redis instance                                    |
| `--redis-prefix`               | `jquery`  | The prefix to assign keys in Redis                               |
| `--redis-ttl`                  | `30m`     | The default TTL for cached responses                             |
| `--redis-connect-timeout`      | `5s`      | The connection timeout for redis                                 |
| `--redis-read-timeout`         | `5s`      | The read timeout for redis                                       |
| `--redis-write-timeout`        | `5s`      | The write timeout for redis                                      |
| `--redis-pool-size`            | `8`       | The number of connection pools to reserve                        |
| `--redis-min-idle-connections` | `0`       | The number of minimum idle connections to maintain               |
| `--log-level`                  | `info`    | The logging level to choose. One of: 'error', 'warning', 'debug' |
| `--log-format`                 | `json`    | The output format to use. One of: 'text', 'json'                 |
| `--log-add-source`             | `false`   | Output the source of the logged entry                            |
| `--proxy-enable`               | `false`   | Whether to enable outbound proxying                              |
| `--proxy-url`                  |           | The URL of the outbound proxy to use                             |
