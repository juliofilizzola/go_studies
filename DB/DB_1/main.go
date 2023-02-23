package main

import (
	"DB_1/db"
	"fmt"
)

func init() {
	fmt.Println("Função init sendo executada")

	db.Connection()
}

func main() {
	fmt.Println("Golang and DB...")

}
