package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	// temos que especificar a quantidade que o array vai ter.
	fmt.Println(array1, "array 1")
	array1[1] = 12
	fmt.Println(array1, "array 1")

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2, "array2")

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3, "array3")
}
