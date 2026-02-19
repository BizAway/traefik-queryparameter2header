# Query Parameter To Header plugin for Traefik

[![Build Status](https://github.com/BizAway/traefik-queryparameter2header/actions/workflows/main.yml/badge.svg)](https://github.com/BizAway/traefik-queryparameter2header/actions/workflows/main.yml)

Middleware plugin for Traefik for setting a request header to the value of a query parameter.


## Configuration

```yaml
queryParameter: "string"  # the source query parameter
header:         "string"  # the destination header
override:       "boolean" # replace the header value if present? (default: true)
```


## Development

Requires `golangci-lint` and `yaegi`.

Use `make` for building, linting and testing.
