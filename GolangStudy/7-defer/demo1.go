package main

import "fmt"

func main1() {
	//defer 压栈入栈
	defer fmt.Println("defer end1")
	defer fmt.Println("defer end2")
	fmt.Println("main:: end")
}
