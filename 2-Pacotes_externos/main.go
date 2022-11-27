package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("aqui é do main")
	erro := checkmail.ValidateFormat("juliofilizzola@hotmail.com")
	fmt.Println(erro)

}
