package main

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resMessages := make([]*Message, 0, 100)
	n := 3 // Number of producers
	var wg sync.WaitGroup
	wg.Add(n)

	inputs := make([]<-chan *Message, n)
	for i := range inputs {
		ch := make(chan *Message)
		inputs[i] = ch
		go simulateProducer(ctx, i, &wg, ch)
	}

	output := make(chan *Message)
	go mergeChannels(ctx, output, inputs...)

	for msg := range output {
		resMessages = append(resMessages, msg)
		log.Printf("Message collected: %v", msg.Data)
	}

	wg.Wait()

	for i, m := range resMessages {
		log.Printf("time of the message #%d = %v", i, m.Time)
	}

	resTimes := make([]int64, 0, len(resMessages))
	for _, m := range resMessages {
		resTimes = append(resTimes, m.Time.UnixMilli())
	}

	isSorted := slices.IsSorted(resTimes)

	log.Printf("Messages are sorted is %v", isSorted)
}

// simulateProducer simulates a producer that sends messages with random times.
func simulateProducer(ctx context.Context, num int, wg *sync.WaitGroup, ch chan<- *Message) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("Producer %d is canceled", num)
			close(ch)
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			data := fmt.Sprintf("Message at %v", time.Now())
			log.Printf("Producer %d writes data: %s", num, data)
			ch <- &Message{Time: time.Now(), Data: data}
		}
	}
}

// mergeChannels merges multiple channels of Messages into a single channel in sorted order.
func mergeChannels(ctx context.Context, output chan<- *Message, inputs ...<-chan *Message) {
	var wg sync.WaitGroup
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	mutex := &sync.Mutex{}

	// This goroutine collects messages from all producers and pushes them into the priority queue.
	for i, ch := range inputs {
		wg.Add(1)
		go func(ch <-chan *Message) {
			defer func() {
				log.Printf("Reading from channel %d is finished", i)
				wg.Done()
			}()
			log.Printf("Start reading from channel %d", i)
			for msg := range ch {
				mutex.Lock()
				heap.Push(&pq, msg)
				mutex.Unlock()
				log.Printf("size of queue = %d", len(pq)) // fixme rm
			}
		}(ch)
	}

	// This goroutine waits for all inputs to finish, then closes the output channel.
	// fixme should be in sync
	go func() {
		wg.Wait()
		log.Println("Closing the output channel")
		close(output)
	}()

	// fixme should be go
	log.Printf("BEFORE Writing to output, len(pq) = %d", len(pq))
	for {
		if len(pq) == 0 {
			continue
		}
		select {
		case <-ctx.Done():
			log.Printf("Merging is canceled") // fixme rm
			return
		case <-time.After(25 * time.Millisecond):
			log.Printf("Flushing to output, len(pq) = %d", len(pq)) // fixme rm
			mutex.Lock()
			for len(pq) > 0 {
				output <- heap.Pop(&pq).(*Message)
			}
			mutex.Unlock()
			log.Printf("Messages flushed") // fixme rm
		}
	}
}

type Message struct {
	Time time.Time
	Data string
}

// PriorityQueue implements heap.Interface and holds Messages.
type PriorityQueue []*Message

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Time.Before(pq[j].Time)
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Message)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
