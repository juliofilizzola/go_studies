package anonymous

import (
	"fmt"
)

func Anonymous() {
	fmt.Println("test")

	func(text string) {
		fmt.Println(text)
	}("Passando o parametro")
}
