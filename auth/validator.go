package auth

import (
	"bytes"
	"encoding/json"
	"github.com/auth0-community/go-auth0"
	"github.com/cstcen/gore/gonfig"
	goreHttp "github.com/cstcen/gore/http"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"io"
	"net/http"
)

func ValidateToken(r *http.Request, extractor auth0.RequestTokenExtractor) (*MemberClaims, error) {
	validator := NewClientValidator(extractor)
	jsonWebToken, err := validator.ValidateRequest(r)
	if err != nil {
		return nil, err
	}
	var claims MemberClaims
	if err := jsonWebToken.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return nil, err
	}
	return &claims, nil
}

func NewClientValidator(extractor auth0.RequestTokenExtractor) *auth0.JWTValidator {
	provider := auth0.NewJWKClient(auth0.JWKClientOptions{
		URI:    gonfig.Instance().GetString("xk5.auth.jwks.url"),
		Client: goreHttp.Instance(),
	}, extractor)
	return NewValidator(provider, extractor)
}

func NewValidator(provider auth0.SecretProvider, extractor auth0.RequestTokenExtractor) *auth0.JWTValidator {
	if provider == nil {
		provider = auth0.NewJWKClient(auth0.JWKClientOptions{
			URI:    gonfig.Instance().GetString("xk5.auth.jwks.url"),
			Client: goreHttp.Instance(),
		}, extractor)
	}
	cfg := auth0.NewConfiguration(provider, []string{}, "", jose.RS256)
	validator := auth0.NewValidator(cfg, extractor)
	return validator
}

func FromParams(r *http.Request) (*jwt.JSONWebToken, error) {
	return auth0.FromParams(r)
}

func FromHeader(r *http.Request) (*jwt.JSONWebToken, error) {
	return auth0.FromHeader(r)
}

func FromBody(r *http.Request) (*jwt.JSONWebToken, error) {
	if r.Body == nil {
		return nil, auth0.ErrTokenNotFound
	}
	var reqBuf bytes.Buffer
	var reqBody []byte
	reqTee := io.TeeReader(r.Body, &reqBuf)
	reqBody, _ = io.ReadAll(reqTee)
	r.Body = io.NopCloser(&reqBuf)
	var body map[string]any
	if err := json.Unmarshal(reqBody, &body); err != nil {
		return nil, err
	}
	access, _ := body["token"].(string)
	if len(access) == 0 {
		access, _ = body["access_token"].(string)
	}
	if len(access) == 0 {
		return nil, auth0.ErrTokenNotFound
	}
	return jwt.ParseSigned(access)
}
