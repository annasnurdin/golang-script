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
	fmt.Printf("%T", args)
}
