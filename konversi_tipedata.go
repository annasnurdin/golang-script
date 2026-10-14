package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // jadi negatif karena kena number overflow
	// int16 itu max nya 32767

	var name = "Annas Nurdin"
	var a = name[0]         // ini byte
	var stringA = string(a) //jadikan string

	fmt.Println(stringA)
}
