package main

import "fmt"

/*
Go语言中大小写区分，类名，属性名，方法名，首字母大写表示对外(其他包)可以访问，否则只能在本包内部访问
*/

type Hero struct {
	Name  string
	Levle int
}

//注意this只是一个变量名，没有其他含义
//*Hero传递指针才能对其更改，否则传进来的是副本
func (this *Hero) Show() {
	fmt.Println("Name is ", this.Name)
	fmt.Println("Level is ", this.Levle)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func main2() {
	hero := Hero{
		Name:  "Jiawen",
		Levle: 10,
	}
	hero.Show()

	hero.SetName("Jiutong")
	hero.Show()

}
