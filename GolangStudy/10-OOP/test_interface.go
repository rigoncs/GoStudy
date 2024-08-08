package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc called...")
	fmt.Println(arg)

	//interface{} 提供区分底层引用的类型是什么
	//类型断言
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	} else {
		fmt.Println("arg is not string type.")
	}
}

type Game struct {
	name string
}

func main() {
	book := Game{"Golang"}
	myFunc(book)
	myFunc(124)
	myFunc("Python")
}
