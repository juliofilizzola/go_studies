package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")
	numberExample := 10
	fmt.Println(numberExample)

	if numberExample > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("menos do que 15")
	}
}