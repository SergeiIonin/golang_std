package timers

import (
	"math/rand"
	"time"
)

func Write(ch chan int, d time.Duration) {

	for {
		// random int
		time.Sleep(d)
		r := rand.Intn(100)
		ch <- r
	}

}