package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save(users user) {
	fmt.Println("save")
	fmt.Println(users)
	fmt.Printf("Salva dada user %s in memory", u.name)
}

func main() {
	fmt.Println("Methods")
	userF := user{"julio", 26}
	userF.save(userF)
}
