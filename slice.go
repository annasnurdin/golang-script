package main

import "fmt"

func main() {
	names := [...]string{"Annas", "Nurdin", "Purwo", "Karin", "Soleh", "Ilma"}
	slice := names[4:6]
	// pointer 4, artinya data pertamanya adalah data index nomor 4 yaitu "Soleh"
	// length nya 2 dari 6-4 = 2, jadi datanya isinya 2 biji
	// capacitynya 2 dari pointer 4 ke data terakhir ada 2 biji yang bisa ditampung

	fmt.Println(names)
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])
}
