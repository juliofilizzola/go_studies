package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var userCreate user
	userCreate.name = "Julio"
	userCreate.age = 26
	fmt.Println(userCreate)

	userCreate2 := user{"Julio", 23}
	fmt.Println(userCreate2)

}
