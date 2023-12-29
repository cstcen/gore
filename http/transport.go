package http

import (
	"bytes"
	"errors"
	"git.tenvine.cn/backend/gore/log"
	"io"
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

	log.DebugCf(ctx, "HTTPClient Req URL:    %q", u.String())
	log.DebugCf(ctx, "HTTPClient Req Header: %q", header)

	if strings.Contains(header.Get("Content-Type"), "application/json") {
		var reqBuf bytes.Buffer
		var reqBody []byte
		if req.Body != nil {
			reqTee := io.TeeReader(req.Body, &reqBuf)
			reqBody, _ = io.ReadAll(reqTee)
			req.Body = io.NopCloser(&reqBuf)
		}
		log.DebugCf(ctx, "HTTPClient Req Body:   %q", reqBody)
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
		log.DebugCf(ctx, "HTTPClient Resp Body:  %q", respBody)
	}

	// Stop timer
	latency := time.Now().Sub(start)
	method := req.Method
	statusCode := resp.StatusCode
	log.DebugCf(ctx, "HTTPClient | %3d | %13v |%-7s %#v",
		statusCode,
		latency,
		method,
		path,
	)

	return resp, nil
}
