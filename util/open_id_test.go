package util

import (
	"testing"
)

func TestGenerateOpenId(t *testing.T) {
	t.Log(GenerateOpenId(EnvSdev0, 1, 10000000000033))
	t.Log(GenerateOpenId(EnvSdev, 1, 10000000000033))
	t.Log(GenerateOpenId(EnvSdev, 2, 10000000000033))
	t.Log(GenerateUnionId(EnvSdev0, 1, 10000000000033))
	t.Log(GenerateUnionId(EnvSdev, 1, 10000000000033))
	t.Log(GenerateUnionId(EnvSdev, 2, 10000000000033))
}
