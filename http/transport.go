package http

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

var (
	ErrUrlNotFound = errors.New("request URL not found")
)

type Transport struct {
	http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	// Start timer
	start := time.Now()
	header := req.Header
	u := req.URL
	if u == nil {
		return nil, ErrUrlNotFound
	}
	path := u.Path
	raw := u.RawQuery
	if len(raw) > 0 {
		path = path + "?" + raw
	}

	slog.DebugContext(ctx, "HTTPClient Request", "URL", u.String())
	slog.DebugContext(ctx, "HTTPClient Request", "Header", header)

	displayBody := strings.Contains(header.Get("Content-Type"), "application/json")
	if displayBody {
		var reqBuf bytes.Buffer
		var reqBody []byte
		if req.Body != nil {
			reqTee := io.TeeReader(req.Body, &reqBuf)
			reqBody, _ = io.ReadAll(reqTee)
			req.Body = io.NopCloser(&reqBuf)
		}
		slog.DebugContext(ctx, "HTTPClient Request", "Body", reqBody)
	}

	resp, err := t.RoundTripper.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		var respBuf bytes.Buffer
		var respBody []byte
		respTee := io.TeeReader(resp.Body, &respBuf)
		respBody, _ = io.ReadAll(respTee)
		resp.Body = io.NopCloser(&respBuf)
		slog.DebugContext(ctx, "HTTPClient Response", "Body", respBody)
	}

	// Stop timer
	latency := time.Now().Sub(start)
	method := req.Method
	statusCode := resp.StatusCode
	slog.DebugContext(ctx, "HTTPClient", "method", method, "uri", path, "status", statusCode, "latency", latency)

	return resp, nil
}
