package main

import "fmt"

func somar(number1 int8, number2 int8) int8 {
	return number1 + number2
}

func main() {
	somar := somar(10, 2)
	fmt.Println(somar)
}
