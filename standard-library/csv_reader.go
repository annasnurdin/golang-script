package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "annas, nurdin\n" +
		"karina, sekar\n" +
		"purwo, fitrianto"

	reader := csv.NewReader(strings.NewReader(csvString))
	// fmt.Println(reader)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
