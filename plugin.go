package queryparameter2header

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	QueryParameter string `json:"queryParameter"`
	Header         string `json:"header"`
	Override       bool   `json:"override"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		QueryParameter: "v",
		Header:         "X-Version",
		Override:       true,
	}
}

// QueryParameterToHeader a QueryParameterToHeader plugin.
type QueryParameterToHeader struct {
	next           http.Handler
	queryParameter string
	header         string
	override       bool
	name           string
}

// New created a new QueryParameterToHeader plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.QueryParameter) < 1 {
		return nil, errors.New("query parameter cannot be empty")
	}
	if len(config.Header) < 1 {
		return nil, errors.New("header cannot be empty")
	}

	return &QueryParameterToHeader{
		header:         config.Header,
		queryParameter: config.QueryParameter,
		override:       config.Override,
		next:           next,
		name:           name,
	}, nil
}

func (a *QueryParameterToHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Has(a.queryParameter) {
		if req.Header.Get(a.header) == "" || a.override {
			req.Header.Set(a.header, strings.Join(req.URL.Query()[a.queryParameter], ","))
		} else {
			_, _ = fmt.Fprintf(os.Stdout, "Header '%s' already present and override set to false, ignoring", a.header)
		}
	}
	a.next.ServeHTTP(rw, req)
}
