package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	//go ping(channel)

	/*msg := ""
	for {
		channel <- "ping"
		msg = <-channel
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
		channel <- "pong"
		msg = <-channel
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}*/

	channel <- "ping"
	for {
		go ping(channel)
		go pong(channel)
	}

	/*go ping(channel)
	msg := <-channel
	fmt.Println("msg = ", msg)*/
}

func ping(c chan string) {
	msg := <-c
	fmt.Println("IN THE PING, MSG = ", msg)
	time.Sleep(500 * time.Millisecond)
	c <- "pong"
}
func pong(c chan string) {
	msg := <-c
	fmt.Println("IN THE PONG, MSG = ", msg)
	time.Sleep(500 * time.Millisecond)
	c <- "ping"
}

/*
func ping(c chan string) {
	msg := <-c
	fmt.Println("pong")
	if msg == "pong" {
		c <- "ping"
	}
}*/
