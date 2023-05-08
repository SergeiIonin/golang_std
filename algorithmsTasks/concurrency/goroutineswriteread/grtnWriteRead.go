package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string, 3)
	//ch := make(chan int, 3)

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("Writing value to channel")
			ch <- fmt.Sprintf("Goroutine %d", i)
		}(i)
	}

	for i := 0; i < 5; i++ {
		go ReceiveFromCh(ch, &wg)
	}

	wg.Wait()

}

func ReceiveFromCh(ch chan string, wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}
