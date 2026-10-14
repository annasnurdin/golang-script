package main

import "fmt"

func main() {
	var nilaiUas = 100
	var nilaiPresensi = 81

	var lulusNilaiUas = nilaiUas > 95
	var lulusPresensi = nilaiPresensi > 80

	var lulus = lulusNilaiUas && lulusPresensi
	fmt.Println(lulus)
}
