package chap3

import "fmt"

func main() {
	const (
		Apple, Banana = iota, iota + 10
		Strawberry, Grape
		Pear, Watermelon
	)

	fmt.Println(Pear)
	fmt.Println(Watermelon)
}
