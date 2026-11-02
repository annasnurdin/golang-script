package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello World\n")
	_, _ = writer.WriteString("hello World\n")
	writer.Flush()

}
