package main

import "fmt"

func deferFunc() int {
	fmt.Println("deferFunc called")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main3() {
	returnAndDefer()
}
