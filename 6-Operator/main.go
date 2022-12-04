package main

import "fmt"

func main() {
	soma := 1 + 2
	sub := 1 - 2
	div := 10 / 2
	mult := 10 * 2
	rest := 10 % 2

	fmt.Println(soma, sub, div, mult, rest)

	const var1 string = "string"
	var2 := "string 2"
	println(var1, var2)

	max := 1 >= 2
	min := 2 <= 1
	igual := 2 == 2

	println(max)
	println(min)
	println(igual)
}
