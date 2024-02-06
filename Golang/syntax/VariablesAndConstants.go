package syntax

import "fmt"

func VariablesAndConstants() {
	var v1 int
	var v2 int
	v1 = 100
	const PI = 3.14
	v3 := 105
	// Automatically assigns the type to V3 based on the value you assigned
	fmt.Println("Value stored in variable V1 :: ", v1)
	fmt.Print("Value stored in variable v2:: ", v2)
	fmt.Print("Value stored in variable v3:: ", v3)
}
