package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")
	timeVar := 0

	for timeVar < 10 {
		timeVar++
		fmt.Println("incrementando")
		time.Sleep(time.Second)
	}

	fmt.Println(timeVar)

	for index := 0; index < 100; index++ {
		fmt.Println("incrementando: ", index)
		time.Sleep(time.Second)
	}
}
