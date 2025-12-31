package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args //ini bentuknya slice
	for _, arg := range args {
		fmt.Println(arg)
	}
	fmt.Printf("%T\n", args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
