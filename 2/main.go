package main

const a = "Hello, World!"

// var b bool
// var c int
var (
	b bool    = true
	c int     = 10
	d string  = "Renan"
	e float64 = 1.2
	f float32
)

func main() {
	var a string = "local"
	println(a) // "local"
	println(b) // false
	b = true
	println(b) // true
	println(c) // 0
	println(d) // ""
	println(e) // +0.000000e+000
	println(f) // +0.000000e+000
	g := "X"   // short syntax
	println(g)
}
