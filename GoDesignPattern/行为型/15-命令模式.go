/*
命令模式的优缺点：
优点：
1.降低系统的耦合度。由于请求者和接收者之间不存在直接引用，因此请求者与接收者之间实现完全解耦，相同的请求者可以对应不同的接收者，同样，相同的接收者也可以供不同的请求者使用，两者之间具有良好的独立性
2.新的命令可以很容易地加入到系统中。由于增加新的具体命令类不会影响到其他类，因此增加新的具体命令类很容易，无需修改原有系统源代码，甚至客户类代码，满足开闭原则
3.可以比较容易地设计一个命令队列或宏命令

缺点：
使用命令模式可能会导致某些系统有过多的具体命令类。因为针对每一个对请求接收者的调用操作都需要设计一个具体命令类，因此在某些系统中可能需要提供大量的具体命令类，这将影响命令模式的使用

适用场景：
1.系统需要将请求调用者和请求接收者解耦，使得调用者和接收者不直接交互。请求调用者无需知道接收者的存在，也无需知道接收者是谁，接收者也无需关心何时被调用。
2.系统需要在不同的时间指定请求、将请求排队和执行请求。一个命令对象和请求的初始调用者可以有不同的生命期，换言之，最初的请求发出者可能已经不在了，而命令对象本身仍然是活动的，可以通过该命令对象去调用请求接收者，而无需关心请求调用者的存在性，可以通过请求日志文件等机制来具体实现
3.系统需要将一组操作组合在一起形成宏命令
*/

package main

import "fmt"

type Cooker struct{}

func (c *Cooker) ToastChicken() {
	fmt.Println("烤鸡肉串")
}

func (c *Cooker) ToastLamp() {
	fmt.Println("烤羊肉串")
}

func (c *Cooker) ToastBeaf() {
	fmt.Println("烤牛排")
}

//抽象的命令
type Command interface {
	Toast()
}

type CommandToastChicken struct {
	c *Cooker
}

func (cmd *CommandToastChicken) Toast() {
	cmd.c.ToastChicken()
}

type CommandToastLamp struct {
	c *Cooker
}

func (cmd *CommandToastLamp) Toast() {
	cmd.c.ToastLamp()
}

type CommandToastBeaf struct {
	c *Cooker
}

func (cmd *CommandToastBeaf) Toast() {
	cmd.c.ToastBeaf()
}

//命令的调用者
type Waiter struct {
	CmdList []Command
}

func (w *Waiter) Notify() {
	if w.CmdList == nil {
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Toast()
	}
}

func main2() {
	cooker := new(Cooker)

	cmdToastChicken := CommandToastChicken{cooker}
	cmdToastLamp := CommandToastLamp{cooker}
	cmdToastBeaf := CommandToastBeaf{cooker}

	waiter := new(Waiter)
	waiter.CmdList = append(waiter.CmdList, &cmdToastBeaf, &cmdToastChicken, &cmdToastLamp)
	waiter.Notify()
}
