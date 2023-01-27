package main

import "fmt"

func invert(nub int) int {
	return nub * -1
}
func invertPoint(nub *int) {
	*nub = *nub * -1
}
func main() {
	number := 20
	numberInvert := invert(number)
	fmt.Println(number, "number original")
	fmt.Println(numberInvert, "number invert")
	newNumber := 40
	fmt.Println(newNumber, "New number")
	invertPoint(&newNumber)
	fmt.Println(newNumber, "new Number invert point")
}
