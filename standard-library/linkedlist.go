package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New() // pointer
	data.PushBack("Annas")
	data.PushBack("Nurdin")
	data.PushBack("Yogyakarta")

	head := data.Front()
	// var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next().Next() // kalau di next terus, bisa nil
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
