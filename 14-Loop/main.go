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
	}

	names := [3]string{"João", "Davi", "Lucas"}
	for index, name := range names {
		// primeiro paramentro do for "of" é o index
		fmt.Println(name, index)
	}

	for _, name := range names {
		fmt.Println(name)
	}

}
