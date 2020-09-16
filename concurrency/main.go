package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	channel = make(chan int, 2)
	num     = 0
	m       sync.Mutex
	wg      sync.WaitGroup
)

func main() {
	go push()
	for num = range channel {
		time.Sleep(3 * time.Second)
		fmt.Printf("got %d from channel\n", num)
	}

	//wg.Add(1)
	//go HelloWorld()
	//wg.Wait()
}

func HelloWorld() {
	println("hello world")
	wg.Done()
}

func push() {
	for i := 0; i < 3; i++ {
		channel <- i
		fmt.Printf("push %d in channel\n", i)
		time.Sleep(1 * time.Second)
	}
	close(channel)
}
