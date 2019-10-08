package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	doSomething()
	waitGroup.Add(1)
	go thinkSomething(&waitGroup)
	waitGroup.Wait()
}

func doSomething() {
	time.Sleep(2 * time.Second)
	fmt.Println("do some thing")
}

func thinkSomething(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("think some thing")
	wg.Done()
}
