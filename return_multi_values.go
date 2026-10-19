package main

import "fmt"

func getFullName() (string, string) {
	return "Annas", "Nurdin"
}

func main() {
	_, lastName := getFullName()
	fmt.Println(lastName)
}
