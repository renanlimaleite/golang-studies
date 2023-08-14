package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 2, 3, 4, 5) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, value := range numeros {
		total += value
	}

	return total
}
