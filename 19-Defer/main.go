package main

import "fmt"

func func1() {
	fmt.Println("Exec func 1")
}

func func2() {
	fmt.Println("Exec Func 2")
}

func main() {
	defer func1()
	func2()
}
