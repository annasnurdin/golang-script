package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Annas"
	// person["address"] = "Yogyakarta"
	person := map[string]string{
		"name":    "Annas",
		"address": "Yogyakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["wrong"]) // key nya tidak ada, default value

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Annas"
	book["wrong"] = "Salah"
	fmt.Println(book)

	delete(book, "wrong") // hapus key
	fmt.Println(book)
}
