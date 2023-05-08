package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(3 * time.Second) // the relevant goroutine will return a value faster than the other
	c <- "Hello from service1!"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service2!"
}

func main() {
	fmt.Println("main started", time.Since(start))

	channel1 := make(chan string)
	channel2 := make(chan string)

	go service1(channel1)
	go service2(channel2)

	// select is like a switch but only for the channels
	select {
	case res := <-channel1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-channel2:
		fmt.Println("Response from service 2", res, time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
