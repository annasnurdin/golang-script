package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked!", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
func main() {
	blacklist := func(name string) bool { // langsung di tulis func nya
		return name == "Anjing"
	}

	registerUser("Annas", blacklist)
	registerUser("Anjing", func(s string) bool { return s == "Anjing" }) // atau bisa langsung
}
