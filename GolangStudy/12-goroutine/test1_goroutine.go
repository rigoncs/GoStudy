package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		fmt.Println("new goroutine : i = ", i)
		time.Sleep(1 * time.Second)
		i++
	}
}

// main goroutine
func main1() {
	go newTask()

	//fmt.Println("main goroutine exit...")

	j := 0
	for {
		fmt.Println("main goroutine : j = ", j)
		time.Sleep(1 * time.Second)
		j++
	}
}
