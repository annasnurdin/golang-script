package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// deklarasi
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		// isi data disini
		data.Value = "Isi data " + strconv.Itoa(i)
		// lalu ambil data selanjutnya
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
