package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // ambil argumen dari go run os.go a b c
	for _, arg := range args {
		fmt.Println(arg)
		// output
		/*
			a
			b
			c
		*/
	}
}
