package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local()) // menggunakan waktu lokal kita
	fmt.Println(now)         // menggunakan waktu lokal di server

	utc := time.Date(2012, time.December, 12, 12, 12, 12, 12, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2012-01-02 12:12:12"

	valueUser := "2020-10-10 01:12:34"

	valueTime, err := time.Parse(formatter, valueUser)
	if err != nil {
		fmt.Println(valueTime)
	} else {
		fmt.Println("Error bro:", err.Error())
	}
}
