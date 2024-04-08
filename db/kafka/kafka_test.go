package kafka

import (
	"github.com/cstcen/gore/gonfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gonfig.Instance().Set("name", "tlogsender")
	gonfig.Instance().Set("env", "sdev0")
	gonfig.Instance().Set("consul", "i-consul-${profile}.xk5.com:8500")
	_ = gonfig.Setup()
}

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{name: "", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConfig()

			assert.NotNil(t, got)
		})
	}
}
