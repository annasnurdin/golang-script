package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Annas", "a"))
	xx := strings.Split("Annas Nurdin", " ")

	fmt.Println(strings.Split("Annas Nurdin", " "))
	fmt.Printf("%T\n", xx)

	fmt.Println(strings.ToLower("Annas Nurdin"))
	fmt.Println(strings.ToUpper("Annas Nurdin"))
	fmt.Println(strings.Replace("Annas Nurdin Annas", "n", "x", 3))
	fmt.Println(strings.ReplaceAll("Annas Nurdin Annas", "n", "x"))
	fmt.Println(strings.Trim("   Annas Nurdin     ", " "))
}
