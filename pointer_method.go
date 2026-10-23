package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	annas := Man{"Annas"}

	annas.Married()

	fmt.Println(annas)
}
