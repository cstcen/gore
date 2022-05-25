package util

import (
	"fmt"
)

func ExampleIntsJoinComma() {

	fmt.Printf("%v", IntsJoinComma([]int{1, 2}))

	// output: 1,2
}

func ExampleUnixToDate() {

	t, err := UnixToDate("1520386162")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t)

	// output: 2018-03-07 09:29:22
}
