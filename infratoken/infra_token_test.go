package infratoken

import (
	"context"
	"github.com/cstcen/gore/gonfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gonfig.Instance().Set("env", "sdev0")
	_ = gonfig.Setup()
}

func TestGet(t *testing.T) {
	type args struct {
		c context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "get", args: args{c: context.Background()}, want: "eyJhbGciOiJIUzI1NiJ9.eyJhZ2VudCI6IlhLNV9TRVJWRVIiLCJhcHBsaWNhdGlvbl9ubyI6Mjk5OTgsImV4cGlyZV90aW1lIjoxNjI4ODQyMjk0NDU1LCJzdnJfaWQiOiJJTkZSQVNFUlZFUjIiLCJlbnYiOiJTREVWMCIsIm1hcmtldF9nYW1lX2lkIjpudWxsLCJ0aW1lc3RhbXAiOjE2Mjg4MjA2OTQ0NTV9.eL6wVWK1uoYddop33r0ylbJohvV1GEgQXwNEdHPjtCs", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
