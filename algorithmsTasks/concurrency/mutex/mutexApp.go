package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	i = i + 1
	wg.Done()
	m.Unlock()
}

func main() {

	i = 0

	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()

	fmt.Println("i = ", i)

}
