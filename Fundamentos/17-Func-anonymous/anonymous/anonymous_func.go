package anonymous

import (
	"fmt"
)

func Anonymous() {
	fmt.Println("Test")

	func(text string) {
		fmt.Println(text)
	}("Passando o parametro")
}
