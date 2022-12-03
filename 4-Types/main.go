package main

import "fmt"

func main() {
	var number1 int64 = 1000000000000000
	fmt.Println(number1)
	var number2 uint64 = 111111111111111
	fmt.Println(number2)

	// alias
	// int32  RUNE

	var number3 rune = 12345
	fmt.Println(number3)

	//uint8 BYTE
	var number4 byte = 123
	fmt.Println(number4)

	// float

	var numberFloat float32 = 123.45

	fmt.Println(numberFloat)

	var numberFloat1 float64 = 1230000000000.45

	fmt.Println(numberFloat1)

	// string

	var str string = "Text"
	fmt.Println(str)

	var char = 'B' // char não existe, ele considera como int

	fmt.Println(char)
}
