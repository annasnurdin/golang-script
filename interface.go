package main

import "fmt"

type HasName interface { // kontrak agar punya apa saja di dalamnya
	GetName() string // harus ada method GetName
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string { // ini implementasi dari interfacenya
	return person.Name
}

func main() {
	person := Person{"Annas"}
	SayHello(person) // baru bisa dipakai function/method nya ini
}
