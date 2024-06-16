package multiple_topics_reader

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
	"testing"
	"time"
)

// todo add test_containers support and ensure test topics exist and have messages before the test runs
func TestKafkaMultipleTopicsReaderStruct_test(t *testing.T) {
	topics := []string{"test_1", "test_2", "test_3"}
	wg := &sync.WaitGroup{}

	reader := MultipleReader{
		Topics:  topics,
		GroupId: fmt.Sprintf("new_%s", time.Now().String()),
	}

	ctx, cancel := context.WithCancel(context.Background())
	count := 0

	go func() {
		select {
		case <-time.After(15 * time.Second):
			t.Logf("15 seconds elapsed")
			cancel()
			return
		}
	}()

	wg.Add(len(topics))
	reader.Read(ctx, wg, &sync.Mutex{}, &count)
	wg.Wait()

	t.Logf("Messages: %d", count)
	if count != 45 {
		t.Fatalf("Expected 45 messages, got %d", count)
	}
}

type MultipleReader struct {
	Topics  []string
	GroupId string
}

func (m *MultipleReader) Read(ctx context.Context, wg *sync.WaitGroup, mut *sync.Mutex, count *int) {
	createReader := func(topic string, groupId string) (reader *kafka.Reader) {
		config := kafka.ReaderConfig{
			Brokers:  []string{"localhost:29092"},
			Topic:    topic,
			GroupID:  groupId, // fixme
			MinBytes: 10e3,    // 10KB
			MaxBytes: 10e6,    // 10MB
		}
		return kafka.NewReader(config)
	}

	read := func(reader *kafka.Reader) {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					log.Printf("Reader is canceled")
					wg.Done()
					return
				}
				wg.Done()
				log.Fatalf("failed to read message: %v", err)
			}
			mut.Lock()
			*count++
			mut.Unlock()
			log.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}

	readers := make([]*kafka.Reader, len(m.Topics))
	for i, topic := range m.Topics {
		readers[i] = createReader(topic, fmt.Sprintf("%s_%d", m.GroupId, i))
		//readers[i] = createReader(topic, fmt.Sprintf("new_%s_%d", time.Now().String(), i))
	}

	for _, reader := range readers {
		go read(reader)
	}
}
