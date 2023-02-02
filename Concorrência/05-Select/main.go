package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select example")
	channel1, channel2 := make(chan string), make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "channel 2"
		}
	}()

	for {
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}

	}
}
