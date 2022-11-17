package http

import (
	"bytes"
	"git.tenvine.cn/backend/gore/log"
	"io"
	"net/http"
	"time"
)

type Transport struct {
	http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	// Start timer
	start := time.Now()
	header := req.Header
	path := req.URL.Path
	raw := req.URL.RawQuery
	if len(raw) > 0 {
		path = path + "?" + raw
	}
	var reqBuf bytes.Buffer
	var reqBody []byte
	var err error
	if req.Body != nil {
		reqTee := io.TeeReader(req.Body, &reqBuf)
		reqBody, _ = io.ReadAll(reqTee)
		req.Body = io.NopCloser(&reqBuf)
	}

	log.DebugCf(ctx, "HTTPClient Req URL:    %s", req.URL.String())
	log.DebugCf(ctx, "HTTPClient Req Header: %+v", header)
	log.DebugCf(ctx, "HTTPClient Req Body:   %+v", string(reqBody))

	resp, err := t.RoundTripper.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	var respBuf bytes.Buffer
	respTee := io.TeeReader(resp.Body, &respBuf)
	respBody, _ := io.ReadAll(respTee)
	resp.Body = io.NopCloser(&respBuf)

	// Stop timer
	latency := time.Now().Sub(start)
	method := req.Method
	statusCode := resp.StatusCode

	log.DebugCf(ctx, "HTTPClient Resp Body:   %+v", string(respBody))

	log.DebugCf(ctx, "HTTPClient | %3d | %13v |%-7s %#v",
		statusCode,
		latency,
		method,
		path,
	)

	return resp, nil
}