package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	LastName string `json:"last_name"`
}

func main() {
	fmt.Println("JSON APPLICATION")
	user1 := user{"Julio", 24, "Filizzola"}

	userJSON, erro := json.Marshal(user1)

	if erro != nil {
		log.Fatal("error")
	}
	fmt.Println(userJSON)
	bytesJSON := bytes.NewBuffer(userJSON)
	fmt.Println(bytesJSON)
}
