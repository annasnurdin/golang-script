package main

import "fmt"

func main() {
	var names [3]string // perlu di deklarasikan jumlah data di dalam array mau berapa biji
	// yang sudah di deklarasikan panjangnya, tidak bisa dihapus lagi

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

	var nilaiTidakPasti = [...]int{ // tapi ini tidak bisa kalau var nilaiTidakPasti [...]int
		1, 2, 3, 4, 5,
	}

	nilaiTidakPasti[0] = 354

	fmt.Println(len(nilaiTidakPasti))
	fmt.Println(nilaiTidakPasti)
}
