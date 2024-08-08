/*
接口隔离原则:
不应该强迫用户的程序依赖他们不需要的接口方法
一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去

合成复用原则：
如果使用继承，会导致父类的任何变换都可能影响到子类的行为。
如果使用对象组合，就降低了这种依赖关系。
对于继承和组合，优先使用组合

迪米特法则：
一个对象应当对其他对象尽可能少的了解，从而降低各个对象之间的耦合，
提高系统的可维护性。例如在一个程序中，各个模块之间相互调用时，通常会提供一个统一的接口来实现。
这样其他模块不需要了解另外一个模块的内部实现细节，这样当一个模块内部的实现发生改变时，不会影响其他模块的使用。(黑盒原理)
*/
package main

import "fmt"

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("cat is eating...")
}

//使用继承,增加sleep方法
type SleepCat struct {
	Cat
}

func (sc *SleepCat) Sleep() {
	fmt.Println("cat is sleeping...")
}

//使用组合
type SleepCatCRP struct {
	c Cat
}

func (sc *SleepCatCRP) Sleep() {
	fmt.Println("cat is sleeping...")
}

func (sc *SleepCatCRP) Eat() {
	sc.c.Eat()
}

func (sc *SleepCatCRP) EatAlt(c Cat) {
	c.Eat()
}

func main4() {
	cat1 := new(Cat)
	cat1.Eat()

	cat2 := new(SleepCat)
	cat2.Sleep()
	cat2.Eat()

	cat3 := new(SleepCatCRP)
	cat3.Eat()
	cat3.Sleep()
}
