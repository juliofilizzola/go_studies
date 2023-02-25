package main

import (
	"fmt"
	"time"
)

func main() {
	channel := white("Hello World!!")
	for index := 0; index < 10; index += 1 {
		fmt.Println(index)
		fmt.Println(<-channel)
	}
}

func white(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("recived value: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
