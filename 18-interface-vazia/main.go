package main

import "fmt"

// interfaces vazias são usadas para definir tipos que não possuem nenhum método
func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da varíavel é %T e valor é %v\n", t, t)
}
