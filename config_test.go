package gocore

import (
	"fmt"
)

type Tes struct {
	A struct {
		B int
		C int
	} `yaml:",flow"`
}

func ExampleUnmarshalConfig() {

	t := new(Tes)
	err := UnmarshalConfig(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", t.A.B)

	// output: 2222
}
