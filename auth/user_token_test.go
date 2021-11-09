package auth

import (
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gonfig.Instance().Set("env", "sdev0")
	gonfig.Instance().Set("name", "gore")
	gonfig.Instance().Set("consul", "i-consul-sdev0.xk5.com:8500")
	gonfig.Setup()
	goreHttp.Setup()
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
		{
			name:    "check",
			args:    args{token: "5e50f76f0b05d3eea4293eb266cdcc6834e161d870d2bf9bcf14872811d9ef4023d62b63bdd27f8e647d1b2a57c54749eecfa13850893545437fdff481dda145b711789d9e8a293128d3a8d41bfd311f7dff5d2a25d5c3f3ef06025ceeea84a3"},
			want:    nil,
			wantErr: false,
		},
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
				assert.Equal(t, "", got.CharacterNo)
			}
		})
	}
}
