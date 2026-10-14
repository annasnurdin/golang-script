package main

import "fmt"

func main() {
	type NoKTP string // membuat alias tipedata baru => NoKTP itu string

	var ktpAnnas NoKTP = "123456789"
	var contoh string = "9876543"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpAnnas)
	fmt.Println(contohKTP)
}
