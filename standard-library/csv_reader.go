package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Annas, Nurdin\n" +
		"Karina, Sekar\n" +
		"Purwo, Fitrianto"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read() // jadi slice
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
