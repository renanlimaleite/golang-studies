package main

// if, switch and select
// não existe else if no golang

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		println("a é maior que b e c é maior que a")
	}

	if a > b || c > a {
		println("a é maior que b ou c é maior que a")
	}

	if a > b {
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println("a é 1")
	case 2:
		println("a é 2")
	case 3:
		println("a é 3")
	default:
		println("default")
	}
}
