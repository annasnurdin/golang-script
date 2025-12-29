package main

import "fmt"

func main(){
	sliceku := []int {1,2,3}
	fmt.Printf("slice = %v\n", sliceku)
	fmt.Println("slice length =",len(sliceku))

	slice2 := append(sliceku, 5, 6)
	fmt.Println("slice 2 =", slice2)

	xx := []int {1,2,3}
	uu := []int {3,4,5,6}

	cc := append(xx, uu...)
	fmt.Println("append 2 slice =", cc)
}