package topics_merger

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
func TestTopicsMerger_test(t *testing.T) {
	topics := []string{"test_1", "test_2", "test_3"}
	merger := Merger{
		Topics:            topics,
		GroupId:           fmt.Sprintf("new_%s", time.Now().String()),
		MergedSourceTopic: "merged_sources",
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
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
	merger.Merge(ctx, wg, &sync.Mutex{}, &count)
	wg.Wait()

	t.Logf("Messages: %d", count)
	if count != 45 {
		t.Fatalf("Expected 45 messages, got %d", count)
	}
}

type Merger struct {
	Topics            []string
	GroupId           string
	MergedSourceTopic string
}

func (merger *Merger) Merge(ctx context.Context, wg *sync.WaitGroup, mut *sync.Mutex, count *int) {
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

	createWriter := func(topic string) (writer *kafka.Writer) {
		return &kafka.Writer{
			Addr:     kafka.TCP("localhost:29092"),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		}
	}

	writer := createWriter(merger.MergedSourceTopic)
	writeMessage := func(writer *kafka.Writer, message kafka.Message) {
		if err := writer.WriteMessages(ctx, message); err != nil {
			log.Fatalf("failed to write message: %v", err)
		}
		mut.Lock()
		*count++
		mut.Unlock()
	}

	readAndWriteToMerged := func(reader *kafka.Reader, writer *kafka.Writer) {
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
			log.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
			m.Headers = append(m.Headers, kafka.Header{
				"initTopic", []byte(m.Topic),
			})
			m.Topic = ""
			go writeMessage(writer, m)
		}
	}

	readers := make([]*kafka.Reader, len(merger.Topics))
	for i, topic := range merger.Topics {
		readers[i] = createReader(topic, fmt.Sprintf("%s_%d", merger.GroupId, i))
	}

	// messages from the same topic will be ordered in the merged topic
	for _, reader := range readers {
		go readAndWriteToMerged(reader, writer)
	}
}
