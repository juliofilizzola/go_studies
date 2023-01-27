package main

import "fmt"

func calMath(n1, n2 int) (soma int, subtract int) {
	soma = n1 + n2
	subtract = n1 - n2
	return
}

func main() {
	fmt.Println("Func returns")

	soma, sub := calMath(1, 2)

	fmt.Println(soma, "soma")
	fmt.Println(sub, "sub")
}
