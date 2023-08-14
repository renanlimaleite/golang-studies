package main

import "fmt"

/*
array, maps, slices
*/

func main() {
	var meuArray [3]int // posições fixas
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[len(meuArray)-1]) // ultimo item

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i, v)
	}
}
