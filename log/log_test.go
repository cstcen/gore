package log

import "fmt"

func ExampleSetupLog() {

	err := SetupLog("debug", "gdis")
	if err != nil {
		fmt.Println(err)
		return
	}

	//std.Info("test1")
	Info("test")

	fmt.Println(std.ReportCaller)

	// output: true
}
