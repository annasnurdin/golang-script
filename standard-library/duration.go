package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Duration ini tipenya int64, nanosecond
	var duration1 time.Duration = 100 * time.Second
	duration2 := 10 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration: %d", duration1/1000000000)
}
