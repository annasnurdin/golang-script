package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	for counter := 1; counter <= 10; counter++ { // cara lain
		fmt.Println("Perulangan ke ", counter)
	}
}
