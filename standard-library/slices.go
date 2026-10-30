package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Annas", "Nurdin", "Purwo", "Karin"}
	fmt.Println(slices.Contains(names, "Purwo"))
	fmt.Println(slices.Index(names, "Karin"))

	values := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(values))
}
