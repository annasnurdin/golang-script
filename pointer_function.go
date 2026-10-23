package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Bantul", "Yogyakarta", ""}
	ChangeAddressToIndonesia(address)

	fmt.Println(address) // ini tidak berubah
	// data struct yang dikirim di function itu dicopy.
}
