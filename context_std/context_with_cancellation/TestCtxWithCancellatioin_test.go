package context_with_cancellation

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// In this case worker goroutine is cancelled after 2 seconds through the cancellation signal and
// then it releases the single waitgroup latch, thus giving control to the main goroutine.
func TestCtxWithCancellatioin_test(t *testing.T) {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	count := 0

	writer(5, ch)
	go worker(ctx, ch, &wg, &count)

	select {
	case <-time.After(2 * time.Second):
		cancel()
	}

	wg.Wait()
	if count != 5 {
		t.Errorf("Expected 5 messages, got %d", count)
	}
}

func worker(ctx context.Context, ch <-chan string, wg *sync.WaitGroup, count *int) {
	for {
		select {
		case msg := <-ch:
			*count += 1
			log.Printf("RECEIVED msg: %s, count = %d", msg, *count)
		case <-ctx.Done():
			log.Printf("Worker cancelled")
			wg.Done()
			return
		}
	}
}

func writer(num int, ch chan<- string) {
	for j := 0; j < num; j++ {
		go func() {
			for i := 0; i < 3; i++ {
				time.Sleep(1 * time.Second)
				msg := fmt.Sprintf("Hello-%d-%d", j, i)
				log.Printf("SENDING msg: %s", msg)
				ch <- msg
			}
		}()
	}
}
