package log

import (
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
)

func init() {
}

func ExampleSetup() {

	gonfig.Instance().Set("gore.logger.level", "info")

	err := Setup()
	if err != nil {
		fmt.Println(err)
		return
	}

	std.Info("test")

	// output:
	// true
}
