package auth

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		ctx   context.Context
		token string
		url   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Member
		wantErr assert.ErrorAssertionFunc
	}{
		{name: "TestCheck", args: args{
			ctx:   context.Background(),
			token: "da240119c5acebe28b4d8b231deabc613f56b58ecae1242eba86175e50a18b87f0178a0b42618f0543eafa1af2db2f3c127ec1fe6bf5e03aa2aa4f134e5bf22ae5f579f9ac8ab7afa7f045f6abdc7650dc691a6d00daa1a395881e5dd8122b6b",
			url:   "https://i-api-sdev0.xk5.com/gateway/verifyToken",
		}, want: &Member{
			Agent:         "XK5_SERVER",
			ApplicationNo: 29998,
			Env:           "sdev0",
			ExpireTime:    1648899010627,
			SvrId:         "INFRASERVER2",
			Timestamp:     1648877410627,
		}, wantErr: assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Check(tt.args.ctx, tt.args.token, tt.args.url)
			if !tt.wantErr(t, err, fmt.Sprintf("Check(%v, %v, %v)", tt.args.ctx, tt.args.token, tt.args.url)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Check(%v, %v, %v)", tt.args.ctx, tt.args.token, tt.args.url)
		})
	}
}
