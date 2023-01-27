package func_var

func Soma(numero ...int) int {
	total := 0
	for _, number2 := range numero {
		total += number2
	}

	return total
}
