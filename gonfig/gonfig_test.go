package gonfig

import (
	"github.com/spf13/viper"
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
			if err := Setup(tt.args.env); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, "0.0.0.0", viper.GetString("messenger.host"))
			assert.Equal(t, "config", viper.GetString("gore.path"))
			assert.Equal(t, "gdis", viper.GetString("gore.cache.appName"))
			assert.Equal(t, "gdis", conf.Gore.Cache.AppName)
		})
	}
}
