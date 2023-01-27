package main

import "fmt"

func closure() func() {
	text := "Dentro da função clousre"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {
	text := "dentro da função main"

	fmt.Println(text)

	funcNew := closure()
	funcNew()
}
