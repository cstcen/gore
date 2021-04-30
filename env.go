package gocore

import (
	"fmt"
	"strings"
)

type Env uint8

const (
	EnvSDev0 Env = iota
	EnvSDev
	EnvDev
	EnvDev2
	EnvDev3
	EnvIOS
	EnvMod
	EnvStg
	EnvXingk5
	EnvXk5
)

var (
	// 当前环境
	EnvCurrent = EnvSDev0

	envCurrentName = "sdev0"

	envMap = map[string]Env{
		"sdev0":  EnvSDev0,
		"sdev":   EnvSDev,
		"dev":    EnvDev,
		"dev2":   EnvDev2,
		"dev3":   EnvDev3,
		"ios":    EnvIOS,
		"mod":    EnvMod,
		"stg":    EnvStg,
		"xingk5": EnvXingk5,
		"xk5":    EnvXk5,
	}
)

func SetupEnv(env string) error {
	e, err := ParseEnv(env)
	if err != nil {
		return err
	}

	envCurrentName = env
	EnvCurrent = e
	return nil
}

func ParseEnv(env string) (Env, error) {
	if e, ok := envMap[strings.ToLower(env)]; ok {
		return e, nil
	}

	var e Env
	return e, fmt.Errorf("invalid env name: %s", env)
}
