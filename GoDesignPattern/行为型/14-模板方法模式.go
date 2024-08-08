/*
模板方法模式：
优点：
1.在父类中形式化地定义一个算法，而由它的子类来实现细节的处理，在子类实现详细的处理算法时并不会改变算法中步骤的执行顺序
2.模板方法模式是一种代码复用技术，它在类库设计中尤为重要，它提取了类库中的公共行为，将公共行为放在父类中，而通过其子类来实现不同的行为，它鼓励我们恰当使用继承来实现代码复用
3.可实现一种反向控制结构，通过子类覆盖父类的钩子方法来决定某一特定步骤是否需要执行
4.在模板方法模式中可以通过子类来覆盖父类的基本方法，不同的子类可以提供基本方法的不同实现，更换和增加新的子类很方便，符合单一职责原则和开闭原则

缺点：
需要为每一个基本方法的不同实现提供一个子类，如果父类中可变的基本方法太多，将会导致类的个数增加，系统更加庞大，设计也更加抽象

适用场景：
1.具有统一的操作步骤或操作过程
2.具有不同的操作细节
3.存在多个具有同样操作步骤的应用场景，但某些具体的操作细节却各不相同。
在抽象类中统一操作步骤，并规定好接口；让子类实现接口，这样可以把各个具体的子类和操作步骤解耦合。
*/
package main

import "fmt"

type Beverage interface {
	Boil() //煮开水
	Brew() //冲泡
	Pour() //倒入杯中
	Add()  //添加小料
}

//封装一套流程,模板基类
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.Boil()
	t.b.Brew()
	t.b.Pour()
	t.b.Add()
}

type MakeCoffee struct {
	template
}

func (mc *MakeCoffee) Boil() {
	fmt.Println("煮水到100℃")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("冲咖啡豆")
}

func (mc *MakeCoffee) Pour() {
	fmt.Println("倒入马克杯中")
}

func (mc *MakeCoffee) Add() {
	fmt.Println("加牛奶和糖")
}

func NewMakeCoffee() *MakeCoffee {
	mc := new(MakeCoffee)
	mc.b = mc
	return mc
}

type MakeTea struct {
	template
}

func (mt *MakeTea) Boil() {
	fmt.Println("煮水到80℃")
}

func (mt *MakeTea) Brew() {
	fmt.Println("冲茶叶")
}

func (mt *MakeTea) Pour() {
	fmt.Println("倒入茶杯中")
}

func (mt *MakeTea) Add() {
	fmt.Println("加柠檬")
}

func NewMakeTea() *MakeTea {
	mt := new(MakeTea)
	mt.b = mt
	return mt
}

func main1() {
	coffeeMaker := NewMakeCoffee()
	coffeeMaker.MakeBeverage()

	teaMaker := NewMakeTea()
	teaMaker.MakeBeverage()
}
