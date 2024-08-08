package main

import (
	"fmt"
)

var gA = 19
var gB float32 = 2.14

func main() {
	var a int = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)

	var b = 3.14
	fmt.Println("b = ", b)
	fmt.Printf("type of b is %T\n", b)

	var c string
	c = "Golang"
	fmt.Println("c = ", c)
	fmt.Printf("type of c is %T\n", c)

	d := 0
	fmt.Println("d = ", d)
	fmt.Printf("type of d is %T\n", d)

	fmt.Println("gA = ", gA, " gB = ", gB)

	var x, y float32 = 1.1, 2.2
	fmt.Println("x = ", x, " y = ", y)

	var (
		xx string  = "avs"
		yy float32 = 3.14
	)
	fmt.Println("xx = ", xx, " yy = ", yy)
}
