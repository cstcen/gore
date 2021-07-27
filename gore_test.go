package gore

import (
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
)

func ExampleSetup() {

	t := new(gonfig.Config)
	err := Setup("sdev0")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", gonfig.GetInstance())

	// output: &{A:{B:2222 C:1111111 D:98} Logger:{Level:debug}}
}
