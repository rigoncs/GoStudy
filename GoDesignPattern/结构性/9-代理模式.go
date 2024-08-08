/*
代理模式优缺点：
优点：
1.能够协调调用者和被调用者，在一定程度上降低了系统的耦合度
2.客户端可以针对抽象主题角色进行编程，增加和更换代理类无须修改源代码，
符合开闭原则，系统具有较好的灵活性和可扩展性

缺点：
代理实现较为复杂
*/
package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

//抽象层
type Shopping interface {
	Buy(goods *Goods)
}

//实现层
type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国购物，买了", goods.Kind)
}

type AmericaShopping struct {
}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国购物，买了", goods.Kind)
}

type OverseasProxy struct {
	shopping Shopping
}

func (op *OverseasProxy) validGoods(good *Goods) bool {
	fmt.Println("对[", good.Kind, "] 进行了辨别真伪")
	if !good.Fact {
		fmt.Println("发现假货,", good.Kind, "不应该购买")
	}
	return good.Fact
}

func (op *OverseasProxy) check(good *Goods) {
	fmt.Println("对[", good.Kind, "]进行了海关检查，成功带回祖国")
}

func (op *OverseasProxy) Buy(good *Goods) {
	//1.先辨别真伪
	if op.validGoods(good) {
		//2.购买货物
		op.shopping.Buy(good) //多态
		//3.海关安检
		op.check(good)
	}
}

func NewProxy(shopping Shopping) OverseasProxy {
	return OverseasProxy{shopping: shopping}
}

//业务逻辑层
func main1() {
	good1 := Goods{
		Kind: "泡菜",
		Fact: true,
	}
	ks := new(KoreaShopping)
	proxy1 := NewProxy(ks)
	proxy1.Buy(&good1)

	good2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}
	as := new(AmericaShopping)
	proxy2 := NewProxy(as)
	proxy2.Buy(&good2)
}
