package main

import "fmt"

func getHelllo(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) { // multiple value
	return "Annas", "Nurdin"
}

func main() {
	result := getHelllo("Annas")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName + " " + lastName)
}
