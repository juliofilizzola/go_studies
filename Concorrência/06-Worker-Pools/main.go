package main

import "fmt"

func main() {
	fmt.Println("Worker-Pools")
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
