package gocore

import "fmt"

func ExampleSetupLog() {

	err := SetupLog("sdev0", "gdis")
	if err != nil {
		fmt.Println(err)
		return
	}

	Info("test")

	fmt.Println(std.ReportCaller)

	// output: true
}
