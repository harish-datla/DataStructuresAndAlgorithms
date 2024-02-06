package syntax

import "fmt"

func BasicDataTypes() {
	var B bool = true
	// Used to store true or false
	var I int = -1000
	// can have positive and negative,
	// other variations of integers are int8, int16, int32 & int64
	var j uint = 1000
	// only stores positive values
	// other variations of integers are uint8, uint16, uint32 & uint64
	var f float32 = 1.2345
	// decimal point number are stored inside floating point variable
	var s string = "Hello, World!"
	// Strings are sequences of Unicode characters. String are immutable.
	var a1 byte = 97 // character a
	// byte are aliaa  of uint8
	var x rune = '\u2615'
	// alias of int32 and are used to store Unicode characters
	var c complex64 = complex(1, 2)
	// store complex number, available in complex64 and complex128
	// complex64 has each component as 32-bit float and
	// complex128  each component has 64-bit float.

	// Automatically assigns the type to V3 based on the value you assigned
	fmt.Println("Value stored in variable B - bool type :: ", B)
	fmt.Println("Value stored in variable I - int type :: ", I)
	fmt.Println("Value stored in variable j - uint type :: ", j)
	fmt.Println("Value stored in variable f - float type :: ", f)
	fmt.Println("Value stored in variable s - string type :: ", s)
	fmt.Println("Value stored in variable a1 - byte type :: ", a1)
	fmt.Printf("Character represented by the byte variable a1 is %c \n", a1)
	fmt.Println("Value stored in variable c - complex type :: ", c)
	fmt.Println("Value stored in variable x - rune type :: ", x)
	fmt.Printf("Character represented by the rune variable x is %c \n ", x)

}
