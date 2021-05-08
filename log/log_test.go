package log

import (
	"fmt"
	"os"
	"strings"
)

func ExampleSetupLog() {

	err := SetupLog()
	if err != nil {
		fmt.Println(err)
		return
	}

	Info("test")

	fmt.Println(std)

	// output: true
}

func ExampleGetRotateLogs() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	dirNames := strings.Split(dir, string(os.PathSeparator))

	fmt.Printf("%+v\n", dir)
	fmt.Printf("%+v\n", dirNames[len(dirNames)-1])

	// output: 1
}
