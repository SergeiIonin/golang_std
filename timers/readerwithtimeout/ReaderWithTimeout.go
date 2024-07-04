package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go addToChan(ch)
	go Run(ch, wg)
	wg.Wait()

}

func addToChan(ch chan<- int) {
	count := 0
	for {
		ch <- count
		count++
		randomNumber := rand.IntN(5)
		fmt.Printf("sleeping for %d seconds\n", randomNumber)
		time.Sleep(time.Duration(randomNumber) * time.Second)
	}
}

// time.After(3 * time.Second) is reset on each iteration
func Run(ch <-chan int, wg *sync.WaitGroup) {
	count := 0
	for {
		select {
		case i := <-ch:
			fmt.Printf("new int = %d \n", i)
		case <-time.After(3 * time.Second):
			fmt.Printf("timeout elapsed!\n")
			count++
			if count == 3 {
				wg.Done()
			}
		}
	}
}
