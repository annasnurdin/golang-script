package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string { // struct validationError punya method Error()
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "annas" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}
func main() {
	err := SaveData("annas", nil)
	if err != nil { // artinya terjadi error
		if validationErr, ok := err.(*validationError); ok { // err.(*validationError) ini type assertion (ngecek tipe datanya)
			// type assertion mengembalikan 2 nilai. validationErr ini kalau datanya benar *validationError
			// ok ini true/false kalau tebakan benar, jadi true dan sebaliknya
			fmt.Println("validation error", validationErr.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("notFound error", notFoundError.Error())
		} else {
			fmt.Println("Unknown error", err.Error())
		}
	} else { // artinya sukses
		fmt.Println("Sukses")
	}

	errr := SaveData("tes", nil)
	if errr != nil {
		switch finalError := errr.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("unknown error", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
