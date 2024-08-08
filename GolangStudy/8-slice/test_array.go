package main

import "fmt"

func printArray1(array [4]int) {
	//形参为固定长度，为值拷贝
	for index, value := range array {
		fmt.Println("index = ", index, ", value = ", value)
	}

	array[0] = 0
}

func main1() {
	//固定长度的数组
	var array1 [10]int
	array2 := [10]int{1, 2, 3, 4}
	array3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	fmt.Printf("array1 type = %T\n", array1)
	fmt.Printf("array2 type = %T\n", array2)
	fmt.Printf("array3 type = %T\n", array3)

	printArray1(array3)
	for index, value := range array3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
