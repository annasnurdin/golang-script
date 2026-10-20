package main

import "fmt"

type Customer struct { // template data/ structure
	Name, Address string
	Age           int
}

func main() {
	var annas Customer
	annas.Name = "Annas"
	annas.Age = 25
	annas.Address = "Yogyakarta"

	fmt.Println(annas)

	purwo := Customer{
		Name:    "Purwo",
		Address: "Magelang",
		Age:     30,
	}

	karin := Customer{"Karin", "Banyuwangi", 25}

	fmt.Println(purwo)
	fmt.Println(karin.Name)
}
