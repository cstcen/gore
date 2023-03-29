package auth

import (
	"bytes"
	"errors"
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/util"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HeaderKeyXRequestDate = "X-Request-Date"
	FormatXRequestDate    = "20060102150405"
)

var (
	ErrIncorrectAuthorization              = errors.New("incorrect authorization")
	ErrIncorrectAuthorizationRequestDate   = errors.New("incorrect authorization request date")
	ErrIncorrectAuthorizationLength        = errors.New("incorrect authorization length")
	ErrIncorrectAuthorizationFirstPart     = errors.New("incorrect authorization first part")
	ErrIncorrectAuthorizationAccess        = errors.New("incorrect authorization access")
	ErrIncorrectAuthorizationSignedHeaders = errors.New("incorrect authorization signed headers")
	ErrIncorrectAuthorizationSignature     = errors.New("incorrect authorization signature")
	ErrAlgorithmNotFound                   = errors.New("algorithm not found")
	ErrAuthConfigNotFound                  = errors.New("auth config not found")
)

type HeaderAuthorization struct {
	Algorithm     string
	Access        string
	SignedHeaders string
	Signature     string
}

// ParseHeaderAuthorization `Authorization` in header
func ParseHeaderAuthorization(authorizationInHeader string) (*HeaderAuthorization, error) {
	authParts := strings.Split(authorizationInHeader, ",")
	if len(authParts) != 3 {
		return nil, fmt.Errorf("%w %v", ErrIncorrectAuthorizationLength, len(authParts))
	}
	authPartFirst := authParts[0]
	authPartSecond := authParts[1]
	authPartThird := authParts[2]
	algorithm, Access, found := strings.Cut(authPartFirst, " ")
	if !found || len(algorithm) == 0 || len(Access) == 0 {
		return nil, ErrIncorrectAuthorizationFirstPart
	}
	accessKey, accessValue, found := strings.Cut(strings.TrimSpace(Access), "=")
	if !found || accessKey != "Access" || len(accessValue) == 0 {
		return nil, ErrIncorrectAuthorizationAccess
	}
	signedHeadersKey, signedHeadersValue, found := strings.Cut(strings.TrimSpace(authPartSecond), "=")
	if !found || signedHeadersKey != "SignedHeaders" || len(signedHeadersValue) == 0 {
		return nil, ErrIncorrectAuthorizationSignedHeaders
	}
	signatureKey, signatureValue, found := strings.Cut(strings.TrimSpace(authPartThird), "=")
	if !found || signatureKey != "Signature" || len(signatureValue) == 0 {
		return nil, ErrIncorrectAuthorizationSignature
	}

	return &HeaderAuthorization{
		Algorithm:     algorithm,
		Access:        accessValue,
		SignedHeaders: signedHeadersValue,
		Signature:     signatureValue,
	}, nil
}

func Authorization(request *http.Request) error {
	authorization := request.Header.Get("Authorization")
	headerAuthorization, err := ParseHeaderAuthorization(authorization)
	if err != nil {
		return err
	}

	secret, err := getAppSecret(headerAuthorization.Access)
	if err != nil {
		return err
	}

	requestDate := getXRequestDate(request.Header)
	if len(requestDate) == 0 {
		return ErrIncorrectAuthorizationRequestDate
	}

	rawQuery := getRawQuery(request.URL)
	payload := getPayload(request)

	var signer util.Signer
	switch util.Algorithm(headerAuthorization.Algorithm) {
	case util.AlgorithmSha256:
		signer = &util.HMacSha256{
			Method:         request.Method,
			Uri:            request.RequestURI,
			RawQuery:       rawQuery,
			Header:         request.Header,
			SignedHeaders:  headerAuthorization.SignedHeaders,
			RequestPayload: payload,
			RequestDate:    requestDate,
			Secret:         secret,
		}
	default:
		return ErrAlgorithmNotFound
	}

	signature, err := signer.Sign()
	if err != nil {
		return err
	}
	if headerAuthorization.Signature != signature {
		return ErrIncorrectAuthorization
	}
	return nil
}

func SetAuthorization(request *http.Request) error {
	requestDate := request.Header.Get(HeaderKeyXRequestDate)
	if len(requestDate) == 0 {
		now := time.Now()
		requestDate = now.Format(FormatXRequestDate)
	}
	request.Header.Set(HeaderKeyXRequestDate, requestDate)
	signedHeaders := make([]string, 0, len(request.Header))
	for k := range request.Header {
		signedHeaders = append(signedHeaders, k)
	}
	rawQuery := getRawQuery(request.URL)
	payload := getPayload(request)
	appKey := gonfig.Instance().GetString("name")
	secret, err := getAppSecret(appKey)
	if err != nil {
		return err
	}
	signer := &util.HMacSha256{
		Method:         request.Method,
		Uri:            request.RequestURI,
		RawQuery:       rawQuery,
		Header:         request.Header,
		SignedHeaders:  strings.Join(signedHeaders, ";"),
		RequestPayload: payload,
		RequestDate:    requestDate,
		Secret:         secret,
	}
	signature, err := signer.Sign()
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s Access=%s, SignedHeaders=%s, Signature=%s", util.AlgorithmSha256, appKey, strings.ToLower(signer.SignedHeaders), signature))
	return nil
}

func getPayload(request *http.Request) []byte {
	if request.Body == nil {
		return nil
	}
	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, request.Body)
	if err != nil {
		return nil
	}
	request.Body = io.NopCloser(&buf)
	return buf.Bytes()
}

func getRawQuery(u *url.URL) string {
	if u == nil {
		return ""
	}
	return u.RawQuery
}

func getAppSecret(appKey string) (string, error) {
	mapString := gonfig.Instance().GetStringMapString("gore.authorization")
	if len(mapString) == 0 {
		return "", ErrAuthConfigNotFound
	}
	secret, ok := mapString[appKey]
	if !ok {
		return "", fmt.Errorf("%s %w", appKey, ErrAuthConfigNotFound)
	}
	return secret, nil
}

func getXRequestDate(header http.Header) string {
	requestDate := header.Get(HeaderKeyXRequestDate)
	if len(requestDate) == 0 {
		return ""
	}
	_, err := time.Parse(FormatXRequestDate, requestDate)
	if err != nil {
		return ""
	}
	return requestDate
}
