package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Hallow\n")
	_, _ = writer.WriteString("Hallow world")
	writer.Flush()
}
