package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("1. Chamando HTTP...")
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()
}
