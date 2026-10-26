package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "annas" {
		return NotFoundError
	}
	//sukses
	return nil
}

func main() {
	err := GetByID("")

	// tidak pakai switch type lagi pakai method dari errors yaitu .Is
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("error validasi")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("error tidak ditemukan")
		} else {
			fmt.Println("error tidak diketahui")
		}
	} else {
		fmt.Println("sukses")
	}
}
