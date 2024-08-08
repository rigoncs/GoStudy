package main

import "fmt"

func foo1(a int, b int) int {
	fmt.Println("=======foo1=======")
	return a + b
}

func foo2(a, b float32) (c, d float32) {
	fmt.Println("========foo2========")
	c = a + b
	d = a - b
	return
}

func foo3(a string, b string) (ret1 string, ret2 string) {
	fmt.Println("=======foo3=======")
	ret1 = a
	ret2 = b
	return
}

func foo4(a float32) (float32, float32) {
	fmt.Println("=========foo4=========")
	return a, a * a
}

func main() {
	ret1 := foo1(10, 10)
	fmt.Println("ret1 = ", ret1)

	ret2, ret3 := foo2(3.14, 3.14)
	fmt.Println("ret2 = ", ret2, " ret3 = ", ret3)

	ret4, ret5 := foo3("Hello, ", "Go!")
	fmt.Println(ret4 + ret5)

	_, ret6 := foo4(2)
	fmt.Println("ret6 = ", ret6)
}
