/*
工厂方法模式的优缺点：
优点：
1.不需要记住具体类名，甚至连具体参数都不用记忆
2.实现了对象创建和使用的分离
3.系统的可拓展性也就变得非常好，无需修改接口和原类
4.对于新产品的创建，符合开闭原则

缺点：
1.增加系统中类的个数，复杂度和理解度增加
2.增加了系统的抽象性和理解难度

适用场景：
1.客户端不知道它所需要的对象的类
2. 抽象工厂类通过其子类来指定创建哪个对象
*/
package pkg1

import "fmt"

// 抽象层
type Shoes interface {
	Show()
}

// 实现层
type XStep struct{}

func (x *XStep) Show() {
	fmt.Println("xtep running!")
}

type LiNing struct{}

func (ln *LiNing) Show() {
	fmt.Println("LiNing running!")
}

// 工厂模块
type ShoesFactory interface {
	CreateShoes() Shoes
}

type XStepFactory struct{}

func (sf *XStepFactory) CreateShoes() Shoes {
	return new(XStep)
}

type LiNingFactory struct{}

func (sf *LiNingFactory) CreateShoes() Shoes {
	return new(LiNing)
}

// 业务逻辑层
func main6() {
	lnFactory := new(LiNingFactory)
	lnShoes := lnFactory.CreateShoes()
	lnShoes.Show()

	xFactory := new(XStepFactory)
	xShoes := xFactory.CreateShoes()
	xShoes.Show()
}
