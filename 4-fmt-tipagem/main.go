package main

import "fmt"

type ID int32

var (
	e float64 = 1.2
	f ID      = 1
)

// %T = tipo
func main() {
	fmt.Printf("O tipo de E é %T, o valor é %v", e, e)
	fmt.Printf("O tipo de F é %T, o valor é %v", f, f)
}
