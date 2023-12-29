package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSgxqSvrTokenString(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "check",
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSgxqSvrTokenString()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSgxqSvrTokenString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("GetSgxqSvrTokenString() got = %v, want %v", got, tt.want)
				return
			}
			claims, err := ParseJwt(got)
			if err != nil {
				t.Errorf("GetSgxqSvrTokenString() parse jwt error = %v", err)
				return
			}
			assert.NotEmpty(t, claims)
			assert.EqualValues(t, "10001", claims.Subject)
		})
	}

}
