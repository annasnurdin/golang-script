package main

import (
	"errors"
	"fmt"
)

func Pembagian(pembilang int, penyebut int) (int, error) {
	if penyebut == 0 {
		return 0, errors.New("Tidak bisa dibagi 0")
	} else {
		return pembilang / penyebut, nil
	}
}
func main() {
	hasil, err := Pembagian(12, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println(err)
	}

	hasil2, err := Pembagian(12, 4)

	if err == nil {
		fmt.Println("Hasil", hasil2)
	} else {
		fmt.Println(err)
	}
}
