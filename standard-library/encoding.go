package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	stringToEncode := "Annas N"

	encoded := base64.StdEncoding.EncodeToString([]byte(stringToEncode))

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error bro: ", err.Error())
	} else {
		fmt.Println(decoded) // ini slice byte ([]byte)
		fmt.Println(string(decoded))
	}
}
