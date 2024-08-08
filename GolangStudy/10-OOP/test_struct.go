package main

import "fmt"

type myInt int

type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	//传递副本，改变不发生变化
	book.title = "Just do it"
}

func changeBook2(book *Book) {
	//传递指针
	book.title = "JUST DO IT"
}

func main1() {
	var a myInt = 10
	fmt.Printf("a = %d, type of a is %T\n", a, a)

	fmt.Println("===========")
	var book1 Book
	book1.title = "<Learn Gollang>"
	book1.auth = "go.org"
	fmt.Printf("%v\n", book1)

	changeBook1(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
