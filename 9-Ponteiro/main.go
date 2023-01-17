package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, "var 1")
	fmt.Println(var2, "var 2")

	var2++

	fmt.Println(var1, "var 1")
	fmt.Println(var2, "var 2")

	// Ponteiro é uma referencia de memória!

}
