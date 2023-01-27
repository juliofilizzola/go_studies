package main

import "fmt"

func func1() {
	fmt.Println("Exec func 1")
}

func func2() {
	fmt.Println("Exec Func 2")
}

func validReciped(n1, n2 float32) bool {
	fmt.Println("validando...")

	media := (n1 + n2) / 2
	fmt.Println(media)
	return !!(media >= 6)
}

func main() {
	defer func1()
	//defe == adiar

	defer fmt.Println(validReciped(23.2, 2))
	func2()
}
