package main

import (
	"container/list"
	"fmt"
)

func main() {
	//list.New ini tipe datanya pointer. pass by reference, ambil datanya dengan .Value
	var data *list.List = list.New()
	//double linked list, bisa nambah depan dan belakang
	data.PushBack("Annas")
	data.PushBack("Nurdin")
	data.PushFront("xx")
	head := data.Front()
	next := head.Next().Value

	fmt.Println(head.Value)
	fmt.Println(next)
	fmt.Println(data)
}
