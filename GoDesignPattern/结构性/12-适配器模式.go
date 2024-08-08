/*
适配器模式：
优点：
1.将目标类和适配者类解耦，通过引入一个适配器类来重用现有的适配者类，无须修改原有结构
2.增加了类的透明性和复用性，将具体的业务实现过程封装在适配者类中，对于客户端类而言是透明的，而且提高了适配者的复用性，同一个适配者类可以在多个不同的系统中复用
3.灵活性和扩展性很好，可以方便地更换适配器，也可以在不修改原有代码的基础上增加新的适配器类，完全符合开闭原则

缺点：
适配器中置换适配者类的某些方法比较麻烦

适用场景：
1.系统需要适用一些现有的类，而这些类的接口不符合系统的需要，甚至没有这些类的源代码
2.想创建一个可以重复使用的类，用于与一些彼此之间没有太大关联的一些类，包括一些可能在将来引进的类一起工作
*/

package main

import "fmt"

//适配的目标
type V5 interface {
	Use5V()
}

//被适配的角色，适配者
type V220 struct {
}

func (v *V220) Use220V() {
	fmt.Println("使用220v的电压")
}

//适配器类
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

//业务类
type Phone struct {
	v5 V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v5: v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone 进行了充电")
	p.v5.Use5V()
}

func main4() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
