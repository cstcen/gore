package http

import (
	"bytes"
	"context"
	"encoding/json"
	"git.tenvine.cn/backend/gore/constant"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/vo"
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

func InternalPost(c context.Context, url, contentType string, body interface{}, expectedPtr interface{}, getInfraToken func(c context.Context) (string, error)) error {
	p, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(p))
	if err != nil {
		return err
	}

	infraToken, err := getInfraToken(c)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", infraToken)
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}

	return nil
}

func Post(url, contentType string, body interface{}, expectedPtr interface{}) error {
	p, err := json.Marshal(body)
	if err != nil {
		return err
	}
	resp, err := cli.Post(url, contentType, bytes.NewReader(p))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func InternalGet(c context.Context, url string, expectedPtr interface{}, getInfraToken func(c context.Context) (string, error)) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	infraToken, err := getInfraToken(c)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", infraToken)
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func Get(url string, expectedPtr interface{}) error {
	resp, err := cli.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func Head(url string, expectedPtr interface{}) error {
	resp, err := cli.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func RespHandler(resp *http.Response, expectedPtr interface{}) error {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyBytes, expectedPtr); err != nil {
		return err
	}
	return nil
}

type Transport struct {
	http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx, ok := req.Context().(context.Context)
	if !ok {
		return nil, vo.BaseResult{Code: http.StatusInternalServerError, Message: "unknown context"}
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
	var reqBody []byte
	var err error
	if req.Body != nil {
		reqTee := io.TeeReader(req.Body, &reqBuf)
		reqBody, _ = io.ReadAll(reqTee)
		req.Body = io.NopCloser(&reqBuf)
	}

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
	}).Info("HTTPClient resp: ")

	return resp, nil
}
