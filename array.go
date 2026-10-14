package main

import "fmt"

func main() {
	var names [3]string // perlu di deklarasikan jumlah data di dalam array mau berapa biji

	names[0] = "Annas"
	names[1] = "Nurdin"
	names[2] = "Purwo"

	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1, 2, 3, // akhirannya tetap koma karena untuk akhiran tetap ada default valuenya
		// int defaultnya 0, string ya string kosong
	}

	var arrayString = [10]string{
		"Annas",
		"Nurdin",
	}

	fmt.Println(values)
	fmt.Println(arrayString)
}
