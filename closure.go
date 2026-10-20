package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++ // memodifikasi variabel di luar func

		inc := 0
		inc++
		fmt.Println(inc)
	}

	increment()
	increment()
	fmt.Println(counter)
}
