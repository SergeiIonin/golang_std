package main

import (
	"GoStudyProject/timers"
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	sleepDur := 4 * time.Second
	base := 1000
	timeoutDur := time.Duration(base) * time.Millisecond // or 1000 * time.Millisecond, but not base * time.Millisecond

	//timeoutDur := time.Duration(500).Milliseconds()

	go timers.Write(ch, sleepDur)

	// in case of timeout 2s and sleepDuration > 2 s, timeout msgs are printed with diff either 2 s (correct) or 4 s (incorrect). WHY?
	for {
		select {
		case v := <-ch:
			fmt.Printf("received value %d\n", v)
		case <-time.After(timeoutDur):
			fmt.Println("timeout reached at ", time.Now().UnixMilli())
		}
	}

}
