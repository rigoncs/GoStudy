/*
策略模式：
优点：
1.提供了对开闭原则的完美支持，用户可以在不修改原有系统的基础上选择算法或行为，也可以灵活地增加新的算法或行为
2.适用策略模式可以避免多重条件选择语句
3.提供了一种算法的复用机制。由于将算法单独提取出来封装在策略类中，因此不同的环境类可以方便地复用这些策略类

缺点：
1.客户端必须知道所有的策略类，并自行决定使用哪一个策略类。只适用于客户端知道所有的算法或行为的情况
2.策略模式将造成系统产生很多具体策略类，任何细小的变化都将导致系统要增加一个新的具体策略类。

适用场景：
准备一组算法，并将每一个算法封装起来，使得它们可以互换
*/

package main

import "fmt"

//抽象策略
type SellStragety interface {
	GetPrice(float32) float32
}

//具体策略
type StragetyA struct {
}

func (a *StragetyA) GetPrice(p float32) float32 {
	return p * 0.8
}

type StragetyB struct{}

func (b *StragetyB) GetPrice(p float32) float32 {
	if p >= 200 {
		p -= 100
	}
	return p
}

type Goods struct {
	Price    float32
	Stragety SellStragety
}

func (g *Goods) SetStragety(s SellStragety) {
	g.Stragety = s
}

func (g *Goods) SellPrice() float32 {
	fmt.Println("原价为", g.Price)
	return g.Stragety.GetPrice(g.Price)
}

func main3() {
	nike := Goods{Price: 200}
	nike.SetStragety(new(StragetyA))
	fmt.Println("上午耐克鞋卖", nike.SellPrice())

	nike.SetStragety(new(StragetyB))
	fmt.Println("下午耐克鞋卖", nike.SellPrice())
}
