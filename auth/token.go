package auth

import (
	"github.com/go-jose/go-jose/v3/jwt"
	"strings"
)

const (
	PrefixBearerAuth = "Bearer "
)

func ParseJwt(tokenString string) (*MemberClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, PrefixBearerAuth)
	token, err := jwt.ParseSigned(tokenString)
	if err != nil {
		return nil, err
	}
	var claims MemberClaims
	if err := token.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return nil, err
	}
	return &claims, nil
}
