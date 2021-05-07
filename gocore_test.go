package gore

import "fmt"

func ExampleSetup() {

	t := new(Tes)
	err := Setup("sdev0", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", t)
	fmt.Printf("%+v\n", conf)

	// output: &{A:{B:2222 C:1111111 D:98} Logger:{Level:debug}}
}
