package gore

import "fmt"

func ExampleSetupLog() {

	err := SetupLog("gdis")
	if err != nil {
		fmt.Println(err)
		return
	}

	Info("test")

	fmt.Println(std.ReportCaller)

	// output: true
}
