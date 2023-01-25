package main

import "fmt"

func recoverExample() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso. ")
	}
}

func valid(n1, n2 float32) bool {
	defer recoverExample()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉIDA É EXATAMENTE 6!")
}

func main() {
	fmt.Println("Panic and Recover")

	fmt.Println(valid(6, 6))

	fmt.Println("Pós funct")
}
