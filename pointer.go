package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1

	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Yogyakarta"

	fmt.Println(address1) // tidak berubah, karena pass by value
	fmt.Println(address2)

}
