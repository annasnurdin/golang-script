package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a[a-z]*s`)

	fmt.Println(regex.MatchString("annas"))
	fmt.Println(regex.MatchString("anas"))
}
