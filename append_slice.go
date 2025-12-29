package main

import "fmt"

func main(){
	sliceku := []int {1,2,3}
	fmt.Printf("slice = %v\n", sliceku)
	fmt.Println("slice length =",len(sliceku))

	slice2 := append(sliceku, 5, 6)
	fmt.Println("slice 2 =", slice2)
}