# jQuery Proxy

A simple web proxy that allows users to perform jQuery operations against an HTTP(S) endpoint with caching.

## Overview

This application provides an OS-agnostic method for performing reliable jQuery expressions against remote HTTP(S) endpoints.

Originally designed to offer an easy method for parsing JSON data for firewalls to obtain IP lists for services published in JSON format, it enables dynamic loading of Alias contents.

Reference the provided documentation for details and examples on how to [deploy](docs/DEPLOYMENT.md) and [use](docs/USAGE.md) this application.

### Details

The application accepts a URL and an xPath query to execute against JSON data retrieved from the URL. The parsed data is returned to the user, enabling JSON filtering without requiring a local package like `jq`.

Redis caching is supported and recommended to prevent too many outbound requests that may lead to rate-limiting.

## License
MIT - Feel free to use, extend, and contribute.