package gocore

import (
	"fmt"
)

func ExampleGetRotateLogs() {

	rotateLogs, err := GetRotateLogs("gdis")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Printf("%+v", rotateLogs)

	// output: 1
}
