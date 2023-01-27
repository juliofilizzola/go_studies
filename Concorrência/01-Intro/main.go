package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("xablau")
	escrever("Hello word")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
