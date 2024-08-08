package main

import (
	"fmt"
	"time"
)

func main2() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("goroutine 结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("向channel中写入", i, " len(c)=", len(c), " cap(c)=", cap(c))
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main goroutine结束")
}
