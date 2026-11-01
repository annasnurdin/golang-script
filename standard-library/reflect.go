package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name, Address, Email string
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with type ", valueField.Type)
	}
	valueType.NumField()
}
func main() {
	readField(Sample{"Annas"})
	readField(Person{"Annas", "Yogyakarta", "annas@annas.com"})
}
