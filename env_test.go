package gocore

import "fmt"

func ExampleSetupEnv() {

	err := SetupEnv("sdev")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(EnvCurrent == EnvSDev)

	// output: true

}
