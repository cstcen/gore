package util

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

type Algorithm string

const (
	AlgorithmSha256 Algorithm = "SHA-256"
)

type Signer interface {
	Sign() (string, error)
}

type HMacSha256 struct {
	Method         string
	Uri            string
	RawQuery       string
	Header         http.Header
	SignedHeaders  string
	RequestPayload []byte

	RequestDate string

	Secret string
}

func (h *HMacSha256) Sign() (string, error) {

	requestStr, err := h.buildRequest()
	if err != nil {
		return "", err
	}
	signing, err := h.buildSigning(requestStr)
	if err != nil {
		return "", err
	}

	return h.buildSignature(h.Secret, signing)
}

func (h *HMacSha256) buildRequest() ([]byte, error) {
	method := h.Method
	uri := h.Uri
	if !strings.HasSuffix(uri, "/") {
		uri = uri + "/"
	}
	rawQuery := h.RawQuery
	header := h.Header
	headerBuf := bytes.Buffer{}
	signedHeaders := strings.Split(h.SignedHeaders, ";")
	for _, k := range signedHeaders {
		v := header.Get(k)
		lowerKey := strings.ToLower(k)
		h := fmt.Sprintf("%s:%s\n", lowerKey, strings.TrimSpace(v))
		headerBuf.WriteString(h)
	}
	headers := headerBuf.String()
	encodedPayload := h.encodePayload()
	return []byte(fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s", method, uri, rawQuery, headers, strings.ToLower(h.SignedHeaders), encodedPayload)), nil
}

func (h *HMacSha256) buildSigning(requestStr []byte) ([]byte, error) {
	hs := crypto.SHA256.New()
	hs.Write(requestStr)
	bs := hs.Sum(nil)
	hashedRequest := strings.ToLower(hex.EncodeToString(bs))
	return []byte(fmt.Sprintf("%s\n%s\n%s", crypto.SHA256.String(), h.RequestDate, hashedRequest)), nil
}

func (h *HMacSha256) buildSignature(secret string, signatureStr []byte) (string, error) {
	hs := hmac.New(sha256.New, []byte(secret))
	hs.Write(signatureStr)
	bs := hs.Sum(nil)
	return hex.EncodeToString(bs), nil
}

func (h *HMacSha256) encodePayload() string {
	if h.RequestPayload == nil {
		return hex.EncodeToString([]byte(""))
	}
	hs := sha256.New()
	hs.Write(h.RequestPayload)
	bs := hs.Sum(nil)
	data := fmt.Sprintf("%x", bs)
	return hex.EncodeToString([]byte(data))
}
