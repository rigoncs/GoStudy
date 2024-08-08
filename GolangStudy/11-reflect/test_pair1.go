/*
变量 可分成 type + value ，这俩构成一个pair
type分为static type和concrete type
static type 为int， string等
concrete type 为interface所指向的具体类型，系统看得见的类型
一个变量要么是static type， 要么是concrete type
*/
package main

import "fmt"

func main1() {
	var a string
	//pair<type:string, value:Golang>
	a = "Golang"

	//pair<type:string, value:Golang>
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
