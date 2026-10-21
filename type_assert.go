package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// var resultInt int = result.(int) => ini kena panic karena tipedatanya tidak sesuai
	// fmt.Println(resultInt)

	hasil := random()
	switch value := hasil.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
