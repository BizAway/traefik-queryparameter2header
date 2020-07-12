package queryparameter2header

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryParameterToHeader(t *testing.T) {
	t.Parallel()

	ctx, handler, recorder := setUp(t)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://test.bizaway.com/home?ap=test1&qp=test2", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Plugin-Test", "test2")
}

func TestQueryParameterToHeaderMultiple(t *testing.T) {
	t.Parallel()

	ctx, handler, recorder := setUp(t)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://test.bizaway.com/home?ap=test1&qp=test2&qp=test3&yap=test4&qp=test5",
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Plugin-Test", "test2,test3,test5")
}

func TestQueryParameterToHeaderOverride(t *testing.T) {
	t.Parallel()

	ctx, handler, recorder := setUp(t)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://test.bizaway.com/home?ap=test1&qp=test2", nil)
	req.Header.Set("X-Plugin-Test", "test?")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Plugin-Test", "test2")
}

func TestQueryParameterToHeaderNoOverride(t *testing.T) {
	t.Parallel()

	cfg := CreateConfig()
	cfg.QueryParameter = "qp"
	cfg.Header = "X-Plugin-Test"
	cfg.Override = false

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://test.bizaway.com/home?ap=test1&qp=test2", nil)
	req.Header.Set("X-Plugin-Test", "test?")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-Plugin-Test", "test?")
}

func setUp(t *testing.T) (context.Context, http.Handler, *httptest.ResponseRecorder) {
	t.Helper()

	cfg := CreateConfig()
	cfg.QueryParameter = "qp"
	cfg.Header = "X-Plugin-Test"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "test")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	return ctx, handler, recorder
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
