package http

import (
	"bytes"
	"context"
	"git.tenvine.cn/backend/gore/constant"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/model"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var cli *http.Client

func GetInstance() *http.Client {
	return cli
}

func init() {
	cli = &http.Client{
		Transport: &Transport{
			RoundTripper: http.DefaultTransport,
		},
		Timeout: constant.TimeoutConn,
	}
}

type Transport struct {
	http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, ok := req.Context().(context.Context)
	if !ok {
		return nil, model.BaseResult{Code: http.StatusInternalServerError, Message: "unknown context"}
	}

	contextLog := log.WithContext(ctx)

	// Start timer
	start := time.Now()
	header := req.Header
	path := req.URL.Path
	raw := req.URL.RawQuery
	if len(raw) > 0 {
		path = path + "?" + raw
	}
	var reqBuf bytes.Buffer
	reqTee := io.TeeReader(req.Body, &reqBuf)
	reqBody, _ := io.ReadAll(reqTee)
	req.Body = io.NopCloser(&reqBuf)

	contextLog.Tracef("HTTPClient url:    %s", req.URL.String())
	contextLog.Tracef("HTTPClient header: %+v", header)
	contextLog.Tracef("HTTPClient body:   %+v", string(reqBody))

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

	contextLog.WithFields(logrus.Fields{
		"method":  method,
		"uri":     path,
		"status":  statusCode,
		"latency": latency,
		"ip":      req.RemoteAddr,
		"body":    string(respBody),
	}).Info("HTTPClient")

	return resp, nil
}
