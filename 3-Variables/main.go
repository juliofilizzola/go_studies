package main

import "fmt"

func main() {
	var (
		var1 = "var 1"
	)
	var2 := "var 2"
	fmt.Printf(var1)
	fmt.Printf(var2)
	var (
		var3 string
		var4 string
	)

	var3 = "Hello"
	var4 = "Stranger"
	fmt.Println(var3)
	fmt.Println(var4)
}
