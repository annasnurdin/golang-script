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

// buat kontrak interface untuk UserSlice nya
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Annas", 21},
		{"Purwo", 22},
		{"Karin", 19},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	test := []User{
		{"Annas", 10},
		{"Purwo", 15},
	}
	sort.Sort(UserSlice(test))
	fmt.Println(test)
}
