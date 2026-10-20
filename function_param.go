package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string { // fungsi yang akan dijadikan parameter
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Annas", filter)
}
