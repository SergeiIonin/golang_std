package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	type Donation struct {
		cond    *sync.Cond
		balance int
	}

	donation := &Donation{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	// Listener goroutines
	f := func(goal int, name string) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
			fmt.Printf("[%s], current value: %d \n", name, donation.balance)
		}
		fmt.Printf("%d goal reached\n", donation.balance)
		donation.cond.L.Unlock()	
	}

	go f(10, "GOAL 10")
	go f(15, "GOAL 15")

	// Updater goroutine
	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.L.Unlock()
		donation.cond.Broadcast()
	}

}
