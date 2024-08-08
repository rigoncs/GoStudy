package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a book.")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book.")
}

func main2() {
	//b pair<type:Book, value: book{}地址>
	b := &Book{} //需要取地址，因为下面赋值给r，interface其实是一个指针

	//r pair<type: value: >
	var r Reader
	//r pair<type:Book, value:book{}地址>
	r = b

	r.ReadBook()

	var w Writer
	//r: pair<type:Book, value:book{}地址>
	w = r.(Writer) //此处断言可以转换成功，因为r此时是Book类型，Book实现了Writer接口

	w.WriteBook()
}
