package auth

import (
	"bytes"
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	AuthorizationInHeader = "SHA-256 Access=123, SignedHeaders=x-request-date;host, Signature=f83753a682e144cd649f5a528ba089b1ea734de27c1e1c5f787a0cb916e0adf2"
)

var (
	request *http.Request
)

func init() {
	req, err := http.NewRequest(http.MethodPost, "/ugc", bytes.NewBufferString("{}"))
	if err != nil {
		panic(err)
	}
	req.Header.Set(HeaderKeyXRequestDate, "20220327160455")
	req.Header.Set("host", "localhost")
	req.Header.Set("Authorization", AuthorizationInHeader)
	request = req

	gonfig.Instance().Set("gore.authorization", map[string]string{"123": "123"})
	gonfig.Instance().Set("name", "123")
}

func TestParseHeaderAuthorization(t *testing.T) {
	type args struct {
		authorizationInHeader string
	}
	tests := []struct {
		name    string
		args    args
		want    *HeaderAuthorization
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "parse",
			args:    args{authorizationInHeader: AuthorizationInHeader},
			want:    &HeaderAuthorization{Algorithm: "SHA-256", Access: "123", SignedHeaders: "host;x-request-date", Signature: "7b74a0ede6c7e3e89d307da6dcb41ce75ba7e29fa145dac7de39751c869f3a2a"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHeaderAuthorization(tt.args.authorizationInHeader)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseHeaderAuthorization(%v)", tt.args.authorizationInHeader)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseHeaderAuthorization(%v)", tt.args.authorizationInHeader)
		})
	}
}

func TestAuthorization(t *testing.T) {
	req := request
	req.Header.Del("Authorization")
	req.Header.Del(HeaderKeyXRequestDate)
	SetAuthorization(req)
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{name: "", args: args{request: req}, wantErr: assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, Authorization(tt.args.request), fmt.Sprintf("Authorization(%v)", tt.args.request))
		})
	}
}

func TestSetAuthorization(t *testing.T) {
	req := request
	req.Header.Del("Authorization")
	type args struct {
		request *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{name: "", args: args{request: req}, wantErr: assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, SetAuthorization(tt.args.request), fmt.Sprintf("SetAuthorization(%v)", tt.args.request))
			assert.Equal(t, AuthorizationInHeader, req.Header.Get("Authorization"))
		})
	}
}
