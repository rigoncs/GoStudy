package main

import "fmt"

func swap(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)
	fmt.Println("a = ", a, " b = ", b)

}
