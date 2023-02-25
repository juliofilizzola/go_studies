package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := multiplexer(write("hello world!"), write("programming in go"))

	for index := 0; index < 10; index += 1 {
		fmt.Println(<-channel)
	}
}

func multiplexer(channel1, channel2 <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			select {
			case message := <-channel1:
				result <- message
			case message := <-channel2:
				result <- message
			}
		}
	}()

	return result
}

func write(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("recive value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
