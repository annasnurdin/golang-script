package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"annas", "nurdin"})
	_ = writer.Write([]string{"karina", "sekar"})
	_ = writer.Write([]string{"purwo", "fitrianto"})

	writer.Flush()
}
