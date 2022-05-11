package log

import (
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
)

func init() {
	gonfig.Instance().Set("gore.logger.filename", "/xk5/logs/gore/gore.log")
	gonfig.Instance().Set("gore.logger.maxage", 1)
}

func ExampleSetup() {

	err := Setup()
	if err != nil {
		fmt.Println(err)
		return
	}

	Info("test")

	fmt.Println(std)

	// output: true
}
