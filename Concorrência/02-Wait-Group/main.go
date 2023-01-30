package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("xablau")
		waitGroup.Done()
	}()

	go func() {
		escrever("Hello word")
		waitGroup.Done()
	}()
	waitGroup.Wait()
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
