package main

import "fmt"

//抽象层
type BeautyWoman interface {
	MakeEyesWithMan()
	MakeLoveWithMan()
}

//实现层
type Pan struct{}

func (p *Pan) MakeEyesWithMan() {
	fmt.Println("潘金莲对本官抛了个媚眼")
}

func (p *Pan) MakeLoveWithMan() {
	fmt.Println("潘金莲与本官共度良宵...")
}

type WangPo struct {
	woman BeautyWoman
}

func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman}
}

func (w *WangPo) MakeEyesWithMan() {
	w.woman.MakeEyesWithMan()
}

func (w *WangPo) MakeLoveWithMan() {
	w.woman.MakeLoveWithMan()
}

//业务逻辑层
func main2() {
	//大官人找金莲，王婆代理
	wangpo := NewProxy(new(Pan))
	//王婆命令金莲
	wangpo.MakeEyesWithMan()
	wangpo.MakeLoveWithMan()
}
