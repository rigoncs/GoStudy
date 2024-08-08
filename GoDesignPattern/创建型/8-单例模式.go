/*
单例模式(饿汉式)：
main函数执行前就分配内存，即使不调用也会分配内存

三个要点：
1.某个类只能有一个实例
2.它必须自行创建这个实例
3.它必须向整个系统提供这个实例

保证一个类永远只能有一个对象，这个对象还能被系统的其他模块使用
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 保证这个类非公有化，外界不能通过这个类直接创建一个对象，首字母小写
type singleton struct{}

/*
饿汉式
*/
// //但是还要有一个指针可以访问这个唯一对象，但是这个指针永远不能改变方向
// // Go没有常指针概念，所以只能通过将指针私有化不让外部模块访问
// var instance *singleton = new(singleton)

// //需要对外提供一个方法来获取到这个对象
// func GetInstance() *singleton {
// 	return instance
// }

/*
懒汉式 与 线程安全
*/
var lock sync.Mutex
var initialized uint32 //标志位

var instance *singleton

func GetInstance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	//如果没有，加锁
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

/*
快捷方式
*/
// var once sync.Once

// var instance *singleton

// func GetInstance() *singleton{
// 	once.Do(func() {
// 		instance = new(singleton)
// 	})
// 	return instance
// }

func (i *singleton) DoSomething() {
	fmt.Println("Something done.")
}

func main8() {
	i1 := GetInstance()
	i1.DoSomething()

	i2 := GetInstance()
	if i1 == i2 {
		fmt.Println("i1 == i2")
	}
}
