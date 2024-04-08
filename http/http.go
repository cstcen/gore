package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var cli = &http.Client{Timeout: 3 * time.Second, Transport: &Transport{
	RoundTripper: http.DefaultTransport,
}}

// GetInstance removed, please use Instance
// Deprecated
func GetInstance() *http.Client {
	return cli
}

func Instance() *http.Client {
	return cli
}

// Setup ...
// Deprecated
func Setup() error {
	return nil
}

type AuthorizationInHeaderGetter interface {
	GetAuthorizationInHeader(req *http.Request) (string, error)
}

type AuthorizationInHeaderSetter interface {
	SetAuthorizationInHeader(request *http.Request) error
}

type AuthorizationInHeaderSetterFunc func(request *http.Request) error

func (fn AuthorizationInHeaderSetterFunc) SetAuthorizationInHeader(request *http.Request) error {
	return fn(request)
}

func InternalPost(ctx context.Context, url, contentType string, body any, expectedPtr any, authorizationInHeaderSetter AuthorizationInHeaderSetter) error {
	r, err := getBodyReader(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, r)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	if err := authorizationInHeaderSetter.SetAuthorizationInHeader(req); err != nil {
		return err
	}

	return Do(req, expectedPtr)
}

func InternalGet(ctx context.Context, url string, expectedPtr any, setAuthorizationInHeader func(request *http.Request) error) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if err := setAuthorizationInHeader(req); err != nil {
		return err
	}
	return Do(req, expectedPtr)
}

func Post(ctx context.Context, url, contentType string, body any, expectedPtr any) error {
	r, err := getBodyReader(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, r)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentType)

	return Do(req, expectedPtr)
}

func Get(ctx context.Context, url string, expectedPtr any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	return Do(req, expectedPtr)
}

func Head(ctx context.Context, url string, expectedPtr any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return err
	}
	return Do(req, expectedPtr)
}

func RespHandler(resp *http.Response, expectedPtr any) error {
	if expectedPtr == nil {
		return nil
	}
	if resp.Body == nil {
		return nil
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(expectedPtr)
}

func Do(req *http.Request, expectedPtr any) error {
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	return RespHandler(resp, expectedPtr)
}

func getBodyReader(body any) (io.Reader, error) {
	if body == nil {
		return nil, nil
	}
	p, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(p), nil
}
