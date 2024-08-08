/*
单一职责原则:
类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个
*/
package main

import "fmt"

type ClothesWork struct {
}

func (cw *ClothesWork) Style() {
	fmt.Println("work style...")
}

type ClothesShop struct {
}

func (cs *ClothesShop) Style() {
	fmt.Println("shop style...")
}

func main1() {
	cw := ClothesWork{}
	cw.Style()

	cs := ClothesShop{}
	cs.Style()
}
