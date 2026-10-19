package main

import "fmt"

func sayHello() {
	fmt.Println("Hello world")
}
func main() {
	sayHello()
	sayHello()
	sayHello()
	sayHello()
	sayName("Annas") // harus dikirim
}

func sayName(name string) { // parameter harus dipakai
	fmt.Println("Hello ", name)
}
