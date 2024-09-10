package main

import (
	"fmt"
	"sync"
)

type single struct {
}

var singleInstance *single
var one sync.Once

func getSingleInstance() *single {
	if singleInstance == nil {
		one.Do(
			func() {
				fmt.Println("Creating a new instance")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Already created a instance")
	}
	return singleInstance
}
func main() {
	for i := 0; i < 10; i++ {
		go getSingleInstance()
	}
	fmt.Scanln()
}
