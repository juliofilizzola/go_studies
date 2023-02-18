package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      uint   `json:"age"`
}

func main() {
	fmt.Println("UnMarshal")
	user1 := `{"name": "julio", "age": 26, "last_name": "Fili"}`
	u := user{}

	if erro := json.Unmarshal([]byte(user1), &u); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(u)
}
