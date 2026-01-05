package main

import (
	"fmt"
	"slices"
)

func main() {
	namanama := []string{"Annas", "Nurdin", "Karin", "Purwo"}
	angkaangka := []int{1, 2, 3, 4, 5, 23, 345, 67, 23, 35324, 14, 234}

	fmt.Println(slices.Min(angkaangka))
	fmt.Println(slices.Max(namanama))

	fmt.Println(slices.Contains(namanama, "annas")) //case sensitive
	fmt.Println(slices.Index(namanama, "Purwo"))

}
