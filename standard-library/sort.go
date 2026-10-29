package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// bikin kontrak interface dulu
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Annas", 25},
		{"Nurdin", 26},
		{"Purwo", 29},
		{"Karin", 20},
	}

	// sort disini, slice user, dijadikan ke UserSlice
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
