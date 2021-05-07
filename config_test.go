package gore

import (
	"fmt"
	"reflect"
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

	s := reflect.ValueOf(t).Elem().FieldByName("Logger").FieldByName("Level").String()

	fmt.Printf("%+v\n", t.A)
	fmt.Printf("%+v\n", t.Logger.Level)
	fmt.Printf("%+v\n", t.Env)
	fmt.Printf("%+v", s)

	// output: {B:2222 C:1111111 D:98}
	// debug
	// sd
	// debug
}
