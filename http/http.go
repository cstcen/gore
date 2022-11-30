package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var cli *http.Client

// GetInstance is replaced by Instance()
// Deprecated
func GetInstance() *http.Client {
	return Instance()
}

func Instance() *http.Client {
	return cli
}

func Setup() error {
	cli = &http.Client{Timeout: 3 * time.Second, Transport: &Transport{
		RoundTripper: http.DefaultTransport,
	}}

	return nil
}

func InternalPost(ctx context.Context, url, contentType string, body any, expectedPtr any, getInfraToken func(c context.Context) (string, error)) error {
	p, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(p))
	if err != nil {
		return err
	}

	infraToken, err := getInfraToken(ctx)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", infraToken)
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}

	return nil
}

func Post(ctx context.Context, url, contentType string, body any, expectedPtr any) error {
	p, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(p))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func InternalGet(ctx context.Context, url string, expectedPtr any, getInfraToken func(c context.Context) (string, error)) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	infraToken, err := getInfraToken(ctx)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", infraToken)
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, url string, expectedPtr any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func Head(ctx context.Context, url string, expectedPtr any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := RespHandler(resp, expectedPtr); err != nil {
		return err
	}
	return nil
}

func RespHandler(resp *http.Response, expectedPtr any) error {
	if expectedPtr == nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bodyBytes, expectedPtr); err != nil {
		return err
	}
	return nil
}
