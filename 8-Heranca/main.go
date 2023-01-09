package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int8
}

type studentes struct {
	person
	course string
}

func main() {
	fmt.Println("Herança")
	p := person{"Julio", "Filizzola", 26}
	p1 := studentes{p, "Ciencia da computação"}
	fmt.Println(p1)
}
