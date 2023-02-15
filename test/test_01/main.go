package main

import (
	"fmt"
	"test_01/address"
)

func main() {
	typeAddress := address.TypeAddress("Rua santos")
	fmt.Println(typeAddress)
}
