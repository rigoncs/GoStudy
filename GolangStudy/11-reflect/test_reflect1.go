/*
reflect包：
ValueOf 获取输入参数接口中的数据的值，如果接口为空则返回0
TypeOf 用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("User is called...")
	fmt.Printf("%v\n", this)
}

func main3() {
	user := User{1, "Golang", 12}

	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过type获取里面的字段
	//1. 获取interface的reflect.Type,通过Type 得到NumField进行遍历
	//2. 得到每个Field，数据类型
	//3. 通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	//通过type获取里面的方法，调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}
