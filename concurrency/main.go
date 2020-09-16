package main

import (
	"fmt"
	"time"
)

var (
	channel = make(chan int, 2)
)

func main() {
	go push()

	for true {
		num, ok := <-channel
		if ok {
			time.Sleep(5 * time.Second)
			fmt.Printf("got %d from channel\n", num)
		} else {
			fmt.Printf("end")
			break
		}
	}

}

func push() {
	for i := 0; i < 6; i++ {
		channel <- i
		fmt.Printf("push %d in channel\n", i)
		time.Sleep(1 * time.Second)
	}
	close(channel)
}
