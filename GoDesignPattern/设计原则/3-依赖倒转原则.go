/*
里氏代换原则：
任何抽象类(interface接口)出现的地方都可以用他的实现类进行替换
实际就是虚拟机制，语言级别实现面向对象功能

依赖倒转原则：
依赖于抽象(接口)，不要依赖于具体的实现(类)
也就是针对接口编程
*/
package main

import "fmt"

//抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//实现层

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW is running...")
}

type Zhang3 struct{}

func (z *Zhang3) Drive(car Car) {
	fmt.Println("Zhang3 is driving...")
	car.Run()
}

type Li4 struct{}

func (l *Li4) Drive(car Car) {
	fmt.Println("Li4 is driving...")
	car.Run()
}

//业务逻辑层
func main3() {
	var benz Car
	benz = new(Benz)

	var zhang3 Driver
	zhang3 = new(Zhang3)
	zhang3.Drive(benz)
}
