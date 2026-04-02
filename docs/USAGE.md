# Usage

## Query Parameters

The parameters supplied to the application must be GET path parameters and must be URL-encoded.

| Name     | Required | Description                                                              |
|----------|----------|--------------------------------------------------------------------------|
| `url`    | Yes      | The target URL from which to pull JSON data for querying.                |
| `query`  | No       | The xPath query to execute against the JSON data.                        |
| `all`    | No       | Determines whether the application should perform `Query` or `QueryAll`. |

> [!NOTE]
> The `url` and `query` parameters *must* be URL-encoded.


## Known Examples

Example usage is provided below for well-known services 

### Google API

**URL**: `https://www.gstatic.com/ipranges/goog.json`

| Description | Usage |
|-------------|-------|
| IPv4 Prefixes | `/?url=https%3A%2F%2Fwww.gstatic.com%2Fipranges%2Fgoog.json&query=%2F%2Fipv4Prefix&all=true` |
| IPv6 Prefixes | `/?url=https%3A%2F%2Fwww.gstatic.com%2Fipranges%2Fgoog.json&query=%2F%2Fipv6Prefix&all=true` |

### GitHub API

**URL**: `https://api.github.com/meta`

| Description | Usage |
|-------------|-------|
| SSH Keys | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=ssh_keys` |
| Hook Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=hooks` |
| Web Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=web` |
| API Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=api` |
| Git Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=git` |
| Importer Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=importer` |
| Enterprise Importer Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=github_enterprise_importer` |
| Packages Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=packages` |
| Pages Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=pages` |
| Actions Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=actions` |
| Mac OS Actions Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=actions_macos` |
| Codespace Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=codespaces` |
| Copilot Prefixes | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=copilot` |
| Website Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Fwebsite` |
| Codespace Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Fcodespaces` |
| Copilot Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Fcopilot` |
| Package Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Fpackages` |
| Actions Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Factions` |
| Inbound Action Full Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Factions_inbound%2Ffull_domains` |
| Inbound Action Wildcard Domains | `/?url=https%3A%2F%2Fapi.github.com%2Fmeta&query=domains%2Factions_inbound%2Fwildcard_domains` |
