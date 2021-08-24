package auth

import (
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gonfig.Instance().Set("env", "sdev0")
	gonfig.Setup()
}

func TestCheckUser(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Member
		wantErr bool
	}{
		{name: "check", args: args{token: "eyJhbGciOiJIUzI1NiJ9.eyJhZ2VudCI6IlhLNV9TRVJWRVIiLCJhcHBsaWNhdGlvbl9ubyI6Mjk5OTgsImV4cGlyZV90aW1lIjoxNjI4ODQyMjk0NDU1LCJzdnJfaWQiOiJJTkZSQVNFUlZFUjIiLCJlbnYiOiJTREVWMCIsIm1hcmtldF9nYW1lX2lkIjpudWxsLCJ0aW1lc3RhbXAiOjE2Mjg4MjA2OTQ0NTV9.eL6wVWK1uoYddop33r0ylbJohvV1GEgQXwNEdHPjtCs"}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUser(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotNil(t, got)
			if got != nil {
				assert.Equal(t, "USER", got.Agent)
			}
		})
	}
}
