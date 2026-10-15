package main

import "fmt"

func main() {
	name := "ANNAS"

	switch name {
	case "ANNAS":
		fmt.Println("Hello ANNAS")
	case "PURWO":
		fmt.Println("Hello PURWO")
	default:
		fmt.Println("Hello, kenalan?")
	}
}
