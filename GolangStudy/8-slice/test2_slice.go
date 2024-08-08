package main

import "fmt"

func main3() {
	slice1 := []int{1, 2, 3}

	var slice2 []int

	var slice3 []int = make([]int, 3)
	slice3[0] = 100

	slice4 := make([]int, 4)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	if slice2 == nil {
		fmt.Println("slice2 is null")
	} else {
		fmt.Println("slice2 is not null")
	}
}
