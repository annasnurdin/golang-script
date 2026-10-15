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

	names := []string{"Annas", "Purwo", "Karin"} // bisa array atau map juga

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
