package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bantul", "Yogyakarta", "Indonesia"}
	address2 := &address1 // < ini pointer
	// penulisan lain: var address2 *Address = &address1

	address2.City = "Kota"

	fmt.Println(address1) // {Kota Yogyakarta Indonesia}
	fmt.Println(address2) // &{Kota Yogyakarta Indonesia}

	*address2 = Address{"Jakarta", "DKI", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}
