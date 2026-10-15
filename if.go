package main

import "fmt"

func main() {
	name := "ANNAS NURDIN"

	if name == "Annas" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Halo, kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
