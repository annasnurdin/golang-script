package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(1997, time.April, 12, 9, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05" //ikuti di dokumentasi

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime) // sudah di format ke Time
	}

	var duration1 time.Duration = time.Second * 1000
	fmt.Println(duration1)

	duration2 := time.Minute * 200
	fmt.Println(duration2)

	duration3 := duration2 - duration1
	fmt.Println(duration3)

	fmt.Printf("duration: %d\n", duration3)
}
