/*
观察者模式:
1.观察者模式可以实现表示层和数据逻辑层的分离，定义了稳定的消息更新传递机制，并抽象了更新接口，使得可以有各种各样不同的表示层充当具体观察者角色
2.观察者模式在观察目标和观察者之间建立一个抽象的耦合。观察目标只需要维持一个抽象观察者的集合，无需了解具体观察者。由于观察目标和观察者没有紧密地耦合在一起，因此它们可以属于不同的抽象化层次
3.观察者模式支持广播通信，观察目标会向所有已注册的观察者对象发送通知，简化了一对多系统设计的难度
4.满足开闭原则的要求，增加新的具体观察者无需修改原有系统代码，在具体观察者与观察目标之间不存在关联关系的情况下，增加新的观察目标也很方便

缺点：
1.如果一个观察目标对象有很多直接和间接观察者，将所有的观察者都通知到会花费很多时间
2.如果在观察者和观察目标之间存在循环依赖，观察目标会触发它们之间进行循环调用，可能导致系统崩溃
3.观察者模式没有相应的机制让观察者知道所观察的目标对象是怎么发生变化的，而仅仅只是知道观察目标发生了变化
*/

package main

import "fmt"

//抽象层
//抽象的观察者
type Observer interface {
	OnTeacherComing()
}

type Notifier interface {
	AddObserver()
	RemoveObserver()
	Notify()
}

//实现层
type StuZhang3 struct {
	Badthing string
}

func (stu *StuZhang3) OnTeacherComing() {
	fmt.Println("张三停止 ", stu.Badthing)
}

type StuLi4 struct {
	Badthing string
}

func (stu *StuLi4) OnTeacherComing() {
	fmt.Println("李四停止 ", stu.Badthing)
}

type StuWang5 struct {
	Badthing string
}

func (stu *StuWang5) OnTeacherComing() {
	fmt.Println("王五停止 ", stu.Badthing)
}

type Moniter struct {
	Observers []Observer
}

func (m *Moniter) AddObserver(observer Observer) {
	m.Observers = append(m.Observers, observer)
}

func (m *Moniter) RemoveObserver(observer Observer) {
	for i, v := range m.Observers {
		if v == observer {
			m.Observers = append(m.Observers[:i], m.Observers[i+1:]...)
			break
		}
	}
}

func (m *Moniter) Notify() {
	for _, v := range m.Observers {
		v.OnTeacherComing()
	}
}

//业务逻辑层
func main() {
	monitor := new(Moniter)
	monitor.AddObserver(&(StuZhang3{Badthing: "玩手机"}))
	monitor.AddObserver(&(StuLi4{Badthing: "抄作业"}))
	monitor.AddObserver(&(StuWang5{Badthing: "看张三玩手机"}))

	monitor.Notify()
}
