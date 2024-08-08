/*
开闭原则：
类的改动是通过增加代码进行的，而不是修改源代码
*/
package main

import "fmt"

type AbstractBanker interface {
	DoBusiness()
}

type TransferBanker struct {
}

func (tb *TransferBanker) DoBusiness() {
	fmt.Println("transfering...")
}

type SaveBanker struct{}

func (sb *SaveBanker) DoBusiness() {
	fmt.Println("saving...")
}

type StockBanker struct{}

func (sb *StockBanker) DoBusiness() {
	fmt.Println("stocking...")
}

//实现一个架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func DoThings(banker AbstractBanker) {
	//通过接口向下来调用(多态)
	banker.DoBusiness()
}

func main2() {
	DoThings(&TransferBanker{})

	DoThings(&SaveBanker{})

	DoThings(&StockBanker{})
}
