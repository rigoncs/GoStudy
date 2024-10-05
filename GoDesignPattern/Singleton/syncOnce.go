package Singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

func getInstanceII() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}
