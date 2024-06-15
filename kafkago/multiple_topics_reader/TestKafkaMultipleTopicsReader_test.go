package multiple_topics_reader

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync"
	"testing"
	"time"
)

// todo add test_containers support and ensure test topics exist and have messages before the test runs
func TestKafkaMultipleTopicsReader_test(t *testing.T) {
	createReader := func(ctx context.Context, topic string, groupId string) (reader *kafka.Reader) {
		config := kafka.ReaderConfig{
			Brokers:  []string{"localhost:29092"},
			Topic:    topic,
			GroupID:  groupId, // fixme
			MinBytes: 10e3,    // 10KB
			MaxBytes: 10e6,    // 10MB
		}
		return kafka.NewReader(config)
	}

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}

	topics := []string{"test_1", "test_2", "test_3"}
	wg.Add(len(topics))

	count := 0

	mut := sync.Mutex{}

	read := func(reader *kafka.Reader, mutex *sync.Mutex) {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					t.Logf("Test finished")
					wg.Done()
					return
				}
				wg.Done()
				t.Fatalf("failed to read message: %v", err)
			}
			mutex.Lock()
			count++
			mutex.Unlock()
			t.Logf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}

	for i, topic := range topics {
		reader := createReader(ctx, topic, fmt.Sprintf("merger_%s_%d", time.Now().String(), i))
		go read(reader, &mut)
	}

	go func() {
		select {
		case <-time.After(15 * time.Second):
			t.Logf("15 seconds elapsed")
			cancel()
			return
		}
	}()

	wg.Wait()

	t.Logf("Messages: %d", count)
	if count != 45 {
		t.Fatalf("Expected 15 messages, got %d", count)
	}
}
