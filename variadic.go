package main

import "fmt"

func sumAll(numbers ...int) int { // ini ...int menandakan bahwa di dalam numbers itu seperti slices
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	hasil := sumAll(1, 2, 3, 4)
	fmt.Println(hasil)

	angkaSlice := []int{10, 10, 10}
	fmt.Println(sumAll(angkaSlice...)) // ditambahkan 3 titik kalau datanya itu dari slice
}
