package auth

import (
	"context"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gonfig.Instance().Set("env", "sdev0")
	gonfig.Instance().Set("name", "banimage")
	gonfig.Instance().Set("consul", "i-consul-sdev0.xk5.com:8500")
	gonfig.Setup()
	goreHttp.Setup()
	goreCache.Setup()
}

func TestExternalCheck(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *Member
		wantErr bool
	}{
		{name: "request", args: args{
			ctx:   context.Background(),
			token: "da240119c5acebe28b4d8b231deabc613f56b58ecae1242eba86175e50a18b87f0178a0b42618f0543eafa1af2db2f3c127ec1fe6bf5e03aa2aa4f134e5bf22ae5f579f9ac8ab7afa7f045f6abdc7650dc691a6d00daa1a395881e5dd8122b6b",
		}, want: nil, wantErr: false},
		{name: "cache", args: args{
			ctx:   context.Background(),
			token: "da240119c5acebe28b4d8b231deabc613f56b58ecae1242eba86175e50a18b87f0178a0b42618f0543eafa1af2db2f3c127ec1fe6bf5e03aa2aa4f134e5bf22ae5f579f9ac8ab7afa7f045f6abdc7650dc691a6d00daa1a395881e5dd8122b6b",
		}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExternalCheck(tt.args.ctx, tt.args.token)

			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}
