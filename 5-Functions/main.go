package main

import "fmt"

func somar(number1 int8, number2 int8) int8 {
	return number1 + number2
}

func calc(n1, n2 int8) (int8, int8) {
	var some = n1 + n2
	var sub = n1 - n2

	return some, sub
}

func main() {
	somar := somar(10, 2)
	fmt.Println(somar)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("text funct")

	fmt.Println(result, "result")

	resSome, resSub := calc(3, 2)

	fmt.Println(resSome)
	fmt.Println(resSub)
}
