package main

import (
	"18-Recursive_functions/fibonacci"
	"fmt"
)

func main() {
	fmt.Println("18-Recursive_functions")
	resolve := fibonacci.Fibonacci(uint(15))
	fmt.Println(resolve)
}
