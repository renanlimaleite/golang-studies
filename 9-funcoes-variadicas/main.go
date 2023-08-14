package main

import (
	"fmt"
)

func main() {
	valor := sum(1, 2, 3, 4) // 10
	fmt.Println(valor)
}

func sum(numeros ...int) int {
	total := 0

	for _, value := range numeros {
		total += value
	}

	return total
}
