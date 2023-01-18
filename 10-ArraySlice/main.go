package main

import (
	"fmt"
	"reflect"
)

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

	slice := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(slice, "slice")

	fmt.Println(reflect.TypeOf(slice), "type of slice")
	fmt.Println(reflect.TypeOf(array3), "type of array")

	slice = append(slice, 00)

	fmt.Println(slice, "new slice")

	fmt.Println("---------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
}
