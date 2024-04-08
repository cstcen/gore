package kafka

import (
	"github.com/cstcen/gore/gonfig"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gonfig.Instance().Set("name", "tlogsender")
	gonfig.Instance().Set("env", "dev3")
	gonfig.Instance().Set("consul", "i-consul-${profile}.xk5.com:8500")
	_ = gonfig.Setup()
}

func TestStartupConsumers(t *testing.T) {
	type args struct {
		handlers map[string]ConsumerMessageHandler
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{handlers: map[string]ConsumerMessageHandler{"member": nil}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StartupConsumers(tt.args.handlers)

			err = ListeningSigterm()

			assert.Nil(t, err)
		})
	}
}
