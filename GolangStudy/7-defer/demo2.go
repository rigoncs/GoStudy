package main

import "fmt"

func defer1() {
	fmt.Println("defer1 called")
}

func defer2() {
	fmt.Println("defer2 called")
}

func defer3() {
	fmt.Println("defer3 called")
}

func main() {
	defer defer1()
	defer2()
	defer defer3()
}
