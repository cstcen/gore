package mongo

import (
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "", args: args{cfg: Config{
			Enable:   true,
			AppName:  "sdev0_gdis",
			Username: "sdev0_gdis_user",
			Password: "sdev0#gdis#ZUrknD",
			Hosts:    []string{"10.251.104.15:27017"},
			Timeout:  3 * time.Second,
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetupDefault(); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
