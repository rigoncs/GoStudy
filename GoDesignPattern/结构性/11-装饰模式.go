/*
装饰模式的优缺点：
优点：
1.对于扩展一个对象的功能，装饰模式比继承更加灵活，不会导致类的个数急剧增加
2.可以通过一种动态的方式来扩展一个对象的功能，从而实现不同的行为
3.可以对一个对象进行多次装饰
4.具体构件类与具体装饰类可以独立变化，用户可以根据需要增加新的具体构件类和具体装饰类，原有类库代码无需改变，符合开闭原则

缺点：
1.使用装饰模式进行系统设计时将产生很多小对象，大量小对象的产生势必会占用更多的系统资源，影响程序的性能。
2.装饰模式提供了一种比继承更加灵活机动的解决方案，但同时也意味着比继承更加易于出错，排错也很困难，对于多次装饰的对象，调试时寻找错误可能需要逐级排查，较为繁琐

适用场景：
1.动态、透明的方式给单个对象添加职责
2.当不能采用继承的方式对系统进行扩展或者采用继承不利于系统扩展和维护时可以使用装饰模式
装饰器模式关注于在一个对象上动态的添加方法，然而代理模式关注于控制对对象的访问。 换句话说，用代理模式，代理类可以对它的客户隐藏一个对象的具体信息。因此，当使用代理模式时，常在一个代理类中创建一个对象的实例。并且，当使用装饰器模式的时候，通常的做法是将原始对象作为一个参数传给装饰者的构造器。
*/
package main

import "fmt"

//抽象层

type Phone interface {
	Show() //构件的功能
}

//抽象的装饰器，装饰器的基础类(该类本应该是interface，但go中的interface语法不可以有成员属性)
type Decorator struct {
	phone Phone //组合进来抽象的phone
}

func (d *Decorator) Show() {}

//实现层
type Huawei struct {
}

func (p *Huawei) Show() {
	fmt.Println("秀出了huawei手机")
}

type Xiaomi struct {
}

func (p *Xiaomi) Show() {
	fmt.Println("秀出了xiaomi手机")
}

//具体的装饰器
type MoDecorator struct {
	Decorator //继承基础的装饰器
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	//装饰器额外的功能
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("带壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

//业务逻辑层
func main3() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.Show()

	fmt.Println("-------")
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()

	fmt.Println("-------")
	var keMoHuawei Phone
	keMoHuawei = NewKeDecorator(moHuawei)
	keMoHuawei.Show()
}
