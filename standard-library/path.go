package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))             // ambil direktorinya
	fmt.Println(path.Base("hello/world.go"))            // ambil nama file nya
	fmt.Println(path.Ext("hello/world.go"))             // ambil extension
	fmt.Println(path.Join("hello", "world", "main.go")) // bikin path full

	fmt.Println(filepath.Join("hello", "world", "main.go")) // filepath untuk OS lokal. di windows pakai \
}
