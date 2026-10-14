package main

import "fmt"

func main() {
	var name string // deklarasi tipe datanya

	name = "Annas Nurdin"
	fmt.Println(name)

	name = "Purwo" // bisa diganti
	fmt.Println(name)

	var umur = "10" // otomatis di berikan tipedatanya
	fmt.Println(umur)

	alamat := "Yogyakarta" // ini sama dengan var alamat = "Yogyakarta"
	fmt.Println(alamat)

	alamat = "Magelang" // ini diubah value variable nya

	var (
		firstName = "Annas"
		lastName  = "Nurdin"
	)

	fmt.Println(firstName + " " + lastName)
}
