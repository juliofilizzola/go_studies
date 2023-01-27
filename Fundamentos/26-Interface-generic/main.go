package main

import "fmt"

func generic(interef interface{}) {
	fmt.Println(interef)
}

func main() {
	generic(32)
	generic("!@")
}
