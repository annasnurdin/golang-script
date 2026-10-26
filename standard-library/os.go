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

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname) // tampil hostname
	} else {
		fmt.Println("Error")
	}
}
