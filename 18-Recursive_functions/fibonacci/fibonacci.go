package fibonacci

import "fmt"

func Fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	fmt.Println(posicao)
	return Fibonacci(posicao-2) + Fibonacci(posicao-1)
}
