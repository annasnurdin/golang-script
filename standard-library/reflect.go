package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "with type ", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
	valueType.NumField()
}
func main() {
	readField(Sample{"Annas"})
	readField(Person{"Annas", "Yogyakarta", "annas@annas.com"})
}
