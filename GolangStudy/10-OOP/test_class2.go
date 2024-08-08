package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human is eating...")
}

func (this *Human) Sleep() {
	fmt.Println("Human is sleeping...")
}

type Superman struct {
	//要继承的类直接写，表示继承
	Human
	level int
}

func (this *Superman) Sleep() {
	fmt.Println("Superman is sleeping...")
}

func (this *Superman) Fly() {
	fmt.Println("Superman is flying...")
}

func main3() {
	human := Human{"zhang3", "male"}
	human.Eat()
	human.Sleep()

	//定义一个子对象
	superman := Superman{Human{"li4", "male"}, 100}
	superman.Eat()
	superman.Sleep()
	superman.Fly()
}
