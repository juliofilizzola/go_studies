package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "hello"
	message := <-channel
	fmt.Println(message)
}
