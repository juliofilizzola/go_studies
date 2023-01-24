package main

import "fmt"

func dayWeek(numberDay int) string {
	switch numberDay {
	case 1:
		return "Doming"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabádo"
	default:
		return "Numero invalido"
	}
}

func main() {
	fmt.Println("Switch")

	resolve := dayWeek(2)

	fmt.Println(resolve)
}
