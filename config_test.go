package gore

import (
	"fmt"
)

type Tes struct {
	A struct {
		B int
		C int
		D int
	} `yaml:",flow"`
	Logger struct {
		Level string
	}
	Env string
}

func ExampleSetupConfig() {

	t := new(Tes)
	err := SetupConfig("sdev0", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", t)

	// output: &{A:{B:2222 C:1111111 D:98} Logger:{Level:debug} Env:sd}
}
