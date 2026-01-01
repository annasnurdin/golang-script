package main

import (
	"fmt"
	"strconv"
)

func main() {
	bool, err := strconv.ParseBool("xx")
	if err != nil {
		fmt.Println("errorr", err.Error())
	} else {
		fmt.Println(bool, "<= ini bool")
	}

	int, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("errorr", err.Error())
	} else {
		fmt.Println(int, "<= ini int")
	}
}
