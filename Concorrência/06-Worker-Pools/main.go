package main

import "fmt"

func main() {
	fmt.Println("Worker-Pools")
	task := make(chan int, 45)
	results := make(chan int, 45)

	go worker(task, results)
	go worker(task, results)
	go worker(task, results)
	go worker(task, results)

	for index := 0; index < 45; index += 1 {
		task <- index
	}

	close(task)

	for index := 0; index < 45; index += 1 {
		result := <-results
		fmt.Println(result)
	}

}

func worker(tasks <-chan int, result chan<- int) {
	for task := range tasks {
		result <- fibonacci(task)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}
