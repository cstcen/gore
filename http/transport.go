package http

import (
	"bytes"
	"git.tenvine.cn/backend/gore/util"
	"io"
	"log"
	"net/http"
	"time"
)

type Transport struct {
	http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	requestId, _ := ctx.Value(util.RequestIDContextKey).(string)

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

	log.Printf("[%s] HTTPClient url:    %s", requestId, req.URL.String())
	log.Printf("[%s] HTTPClient header: %+v", requestId, header)
	log.Printf("[%s] HTTPClient body:   %+v", requestId, string(reqBody))

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

	log.Printf("[%s] HTTPClient method=%s uri=%s status=%v latency=%v ip=%s body=%s",
		requestId,
		method,
		path,
		statusCode,
		latency,
		req.RemoteAddr,
		string(respBody),
	)

	return resp, nil
}
