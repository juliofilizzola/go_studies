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

	// Ponteiro é uma referência de memória!
	var var3 int
	var point *int
	var3 = 100
	/**
	point = var3
	desse jeito o go não deixar atribuir, aqui você terar
	que add o & antes do valor que será atribuido
	*/
	point = &var3

	fmt.Println(var3, "var 3")
	fmt.Println(point, "point")        // aqui vai imprimir o valor do endereço da memória
	fmt.Println(*point, "point value") // com o * consegue recuperar o valor armazenado
	var3 = 1233
	fmt.Println(*point, "point value")
	fmt.Println(var3, "var 3")

}
