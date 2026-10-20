package main

import "fmt"

type Customer struct { // template data/ structure
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) { // membuat struct method
	fmt.Println("Hello,", name, "my name is", customer.Name)
}

func main() {
	var annas Customer
	annas.Name = "Annas"
	annas.Age = 25
	annas.Address = "Yogyakarta"

	fmt.Println(annas)

	purwo := Customer{ // objek purwo
		Name:    "Purwo",
		Address: "Magelang",
		Age:     30,
	}

	karin := Customer{"Karin", "Banyuwangi", 25}

	fmt.Println(purwo)
	fmt.Println(karin.Name)

	//sayHello() ini tidak bisa karena method dari struct Customer
	purwo.sayHello("Soleh") // panggil method dari struct
}
