package http

import (
	"bytes"
	"context"
	"encoding/json"
	"git.tenvine.cn/backend/gore/constant"
	"io"
	"net/http"
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
