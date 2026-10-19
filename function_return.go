package main

import "fmt"

func getHelllo(name string) string {
	return "Hello " + name
}

func main() {
	result := getHelllo("Annas")
	fmt.Println(result)
}
