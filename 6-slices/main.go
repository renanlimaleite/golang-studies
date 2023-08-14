package main

import "fmt"

/*
slices
*/

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=10 cap=10 [10 20 30 40 50 60 70 80 90 100]

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // len=0 cap=10 []

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // len=4 cap=10 [10 20 30 40]

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // len=8 cap=8 [30 40 50 60 70 80 90 100]

	s = append(s, 200)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // len=9 cap=18 [30 40 50 60 70 80 90 100 200]

}
