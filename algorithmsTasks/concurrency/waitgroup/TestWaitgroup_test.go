package waitgroup

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

// Writer spawns n goroutines, each of which sends 3 messages to the channel ch.
// Waitgroup is utilized to add 3 quotas for each goroutine and reader goroutine takes quota back upon reading
// with wg.Add(). In this case reader sync is guaranteed bc reader's channel is blocked msg is sent to chan
// and additional sync by mutex is not needed.

func TestCtxWithCancellatioin_test(t *testing.T) {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	writer(5, ch, &wg)
	count := 0
	go reader(ch, &wg, &count)
	wg.Wait()
	if count != 15 {
		t.Errorf("Expected 15 messages, got %d", count)
	}
}

func writer(num int, ch chan<- string, wg *sync.WaitGroup) {
	for j := 0; j < num; j++ {
		wg.Add(3)
		go func() {
			for i := 0; i < 3; i++ {
				msg := fmt.Sprintf("Hello-%d-%d", j, i)
				log.Printf("SENDING msg: %s", msg)
				ch <- msg
			}
		}()
	}
}

func reader(ch <-chan string, wg *sync.WaitGroup, count *int) {
	for {
		select {
		case msg := <-ch:
			*count += 1
			wg.Done()
			log.Printf("%s, count = %d", msg, *count)
		}
	}
}
