package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Annas", "Nurdin"})
	_ = writer.Write([]string{"Purwo", "Fitrianto"})
	_ = writer.Write([]string{"Karina", "Sekar"})

	writer.Flush()
}
