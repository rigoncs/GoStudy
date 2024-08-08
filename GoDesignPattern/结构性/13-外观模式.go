/*
外观模式：
优点：
1.对客户端屏蔽了子系统组件，减少了客户端所需处理的对象数目，并使得子系统使用起来更加容易。通过引入外观模式，客户端代码将变得很简单，与之关联的对象也很少
2.实现了子系统与客户端之间的松耦合关系，使得子系统的变化不会影响到调用它的客户端，只需要调整外观类即可
3.一个子系统的修改对其他子系统没有任何影响

缺点：
1.不能很好地限制客户端直接使用子系统类，如果对客户端访问子系统类做太多的限制则减少了可变性和灵活性
2.如果设计不当，增加新的子系统可能需要修改外观类的源代码，违背了开闭原则

适用场景：
1.复杂系统需要简单入口使用
2.客户端程序与多个子系统之间存在很大的依赖性
3.在层次化结构中，可以使用外观模式定义系统中每一层的入口，层与层之间不直接产生联系，而通过外观类建立联系，降低层之间的耦合度
*/

package main

import "fmt"

type SubSystemA struct{}

func (a *SubSystemA) MethodA() {
	fmt.Println("subsystem A do something")
}

type SubSystemB struct{}

func (b *SubSystemB) MethodB() {
	fmt.Println("subsystem B do something")
}

type SubSystemC struct{}

func (c *SubSystemC) MethodC() {
	fmt.Println("subsystem C do something")
}

type System struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
}

func (s *System) Task1() {
	s.a.MethodA()
	s.b.MethodB()
}

func (s *System) Task2() {
	s.a.MethodA()
	s.c.MethodC()
}

func main() {
	s := System{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
	}

	s.Task1()
	s.Task2()
}
