package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5) //pointer
	// kalau valuenya ada yang kosong, maka dia nil
	// datanya akan mulai dari paling terakhir

	data.Value = "Value 1"

	data = data.Next()
	data.Value = "Value 2"

	data = data.Next()
	data.Value = "Value 3"

	data = data.Next()
	data.Value = "Value 4"

	data = data.Next()
	data.Value = "Value 5"

	fmt.Println(data.Value)
	fmt.Println("============")

	data.Do(func(a any) {
		fmt.Println(a)
	})
	fmt.Println("===============")

	data2 := ring.New(10)
	for i := 0; i < data2.Len(); i++ {
		data2.Value = "Value" + strconv.Itoa(i)
		data2 = data2.Next()
	}

	data2.Do(func(a any) {
		fmt.Println(a)
	})

}
