package main

import (
	"fmt"
	"time"
)

func main() {

	//创建一个形参为空，返回值为空的匿名函数
	func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	res := func(a int, b int) bool {
		fmt.Println("a = ", a, " b = ", b)
		return a == b
	}(10, 20)
	println("res = ", res)

	for {
		time.Sleep(1 * time.Second)
	}
}
