package main

import "fmt"

type HasName interface { // kontrak agar punya apa saja di dalamnya
	GetName() string // harus ada method GetName
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func main() {
	SayHello(???)
}
