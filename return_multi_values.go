package main

import "fmt"

func getFullName() (string, string) {
	return "Annas", "Nurdin"
}
func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
