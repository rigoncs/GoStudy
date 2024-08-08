package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)

	slice1 = append(slice1, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice1), cap(slice1), slice1)

	fmt.Println("============")
	slice2 := make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice2), cap(slice2), slice2)
	slice2 = append(slice2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice2), cap(slice2), slice2)

	slice3 := []int{1, 2, 3}
	slice4 := slice3[:2]
	fmt.Println(slice4)

	slice4[0] = 99
	fmt.Println(slice3)
	fmt.Println(slice4)

	slice5 := make([]int, 3)
	copy(slice5, slice3)
	slice5[1] = 66
	fmt.Println(slice3)
	fmt.Println(slice5)
}
