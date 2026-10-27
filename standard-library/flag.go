package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Isi database host")

	flag.Parse()
	fmt.Println(host)  // reference memori nya
	fmt.Println(*host) // ini valuenya
}
