package main

import "fmt"

func main() {
	names := [...]string{"Annas", "Nurdin", "Purwo", "Karin", "Soleh", "Ilma"}
	slice := names[4:6]
	// pointer 4, artinya data pertamanya adalah data index nomor 4 yaitu "Soleh"
	// length nya 2 dari 6-4 = 2, jadi datanya isinya 2 biji
	// capacitynya 2 dari pointer 4 ke data terakhir ada 2 biji yang bisa ditampung
	slice2 := names[:3] // dari 0 sampai index 3 -1
	slice3 := names[3:] // dari index ke 3 sampai terakhir
	slice4 := names[:]

	fmt.Println(names)
	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"} // ini array

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu baru" // ini mengganti value di slice sekaligus di array nya
	daySlice1[1] = "Minggu baru"
	fmt.Println(days)
	fmt.Println(daySlice1)

	daySlice2 := append(daySlice1, "Libur baru") // ini membuat slice baru dari slice daySlice1
	fmt.Println(daySlice2)
	fmt.Println(days) // tetapi arraynya masih sama karena panjangnya melebihi

}
