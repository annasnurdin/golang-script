package main

import "fmt"

func main() {
	var names [3]string // perlu di deklarasikan jumlah data di dalam array mau berapa biji

	names[0] = "Annas"
	names[1] = "Nurdin"
	names[2] = "Purwo"

	fmt.Println(names[1])
	fmt.Println(names[2])
}
