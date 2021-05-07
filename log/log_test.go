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

	//std.Info("test1")
	Info("test")

	fmt.Println(std.ReportCaller)

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
