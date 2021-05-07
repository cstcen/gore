package gocore

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
}

func ExampleUnmarshalConfig() {

	t := new(Tes)
	err := UnmarshalConfig("sdev0", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", t.A)
	fmt.Printf("%+v\n", t.Logger.Level)

	// output: {B:2222 C:1111111 D:98}
	// debug
}
