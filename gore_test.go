package gore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Viper().Set("env", "sdev0")
}

func TestSetup(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "setup", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setup(); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.NotNil(t, Cache())
			assert.NotNil(t, Redis())
		})
	}
}
