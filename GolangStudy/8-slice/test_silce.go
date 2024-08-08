package main

import "fmt"

func printArray(array []int) {
	//动态数组，引用传递
	for _, value := range array {
		fmt.Println("value = ", value)
	}
	array[0] = 0
}

func main2() {
	array1 := []int{1, 2, 3} //动态数组，切片
	fmt.Printf("array1 type is %T\n", array1)
	printArray(array1)
	for _, value := range array1 {
		fmt.Println("value = ", value)
	}
}
