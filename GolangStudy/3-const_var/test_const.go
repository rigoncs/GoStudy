package main

import "fmt"

//利用const实现枚举
const (
	BEIJING = iota
	SHANGHAI
	SHENZHEN

	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 10, iota * 20
	i, j
)

func main() {

	const pi float32 = 3.14
	fmt.Println("pi = ", pi)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, " b = ", b)
	fmt.Println("a = ", c, " b = ", d)
	fmt.Println("a = ", e, " b = ", f)
	fmt.Println("a = ", g, " b = ", h)
	fmt.Println("a = ", i, " b = ", j)
}
