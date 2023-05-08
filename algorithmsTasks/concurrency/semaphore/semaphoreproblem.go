package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	/*channel <- red
	msg := <-channel
	fmt.Println("msg = ", msg)*/

	/*go semaphore(channel)
	msg := <-channel
	fmt.Println("msg = ", msg)*/

	/*msg := ""
	for {
		channel <- red
		msg = <-channel
		fmt.Println(msg)
		time.Sleep(700 * time.Millisecond)

		channel <- yellow
		msg = <-channel
		fmt.Println(msg)
		time.Sleep(300 * time.Millisecond)

		channel <- green
		msg = <-channel
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}*/

	start(channel)
	go semaphore(channel)

	/*for {
		listener(channel)
	}*/

	//go start(channel)

	defer afterAll()
}

const red = "red"
const yellow = "yellow"
const green = "green"

func start(c chan string) {
	c <- red
}

func listener(c chan string) {
	msg := <-c
	switch msg {
	case "red":
		fmt.Println("RED")
		time.Sleep(700 * time.Millisecond)
		//fmt.Println("SENDING YELLOW")
		c <- yellow
	case "yellow":
		fmt.Println("YELLOW")
		time.Sleep(300 * time.Millisecond)
		c <- green
	case "green":
		fmt.Println("GREEN")
		time.Sleep(500 * time.Millisecond)
		c <- red
	}
}

func semaphore(c chan string) {
	for {
		msg := <-c
		switch msg {
		case "red":
			fmt.Println("RED")
			time.Sleep(700 * time.Millisecond)
			//fmt.Println("SENDING YELLOW")
			c <- yellow
		case "yellow":
			fmt.Println("YELLOW")
			time.Sleep(300 * time.Millisecond)
			c <- green
		case "green":
			fmt.Println("GREEN")
			time.Sleep(500 * time.Millisecond)
			c <- red
		default:
			panic("unknown command")
		}
	}
}

func afterAll() {
	fmt.Println("Semaphore is out of service")
}
