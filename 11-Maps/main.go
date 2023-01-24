package main

import "fmt"

func main() {
	fmt.Println("Maps")
	user := map[string]string{
		"name":     "Julio",
		"lastName": "Filizzola",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"fist": "Julio",
			"last": "Filizzola",
		},
		"Learn": {
			"name": "Cience",
		},
	}

	fmt.Println(user2)
}
