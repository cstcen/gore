package gonfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {
	type args struct {
		env    string
		outPtr []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "setup", args: args{env: "sdev0"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setup(tt.args.env, "gore"); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, false, vp.GetBool("gore.cache.enable"))
			//assert.Equal(t, "0.0.0.0", vp.GetString("messenger.host"))
			//assert.Equal(t, "config", vp.GetString("gore.path"))
			assert.Equal(t, "gdis", vp.GetString("gore.cache.appName"))
			//assert.Equal(t, "gdis", conf.Gore.Cache.AppName)
		})
	}
}
