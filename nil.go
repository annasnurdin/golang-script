package main

import "fmt"

// nil hanya untuk:
/*
interface
function
map
slice
pointer
channel
*/

// func Contoh(name string) string { ini tidak bisa karena hanya untuk tipe data di atas (ini string)
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(data)
	}

}
