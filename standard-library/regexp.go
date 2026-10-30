package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-zA-Z]+)s`)

	fmt.Println(regex.MatchString("annas"))
}
