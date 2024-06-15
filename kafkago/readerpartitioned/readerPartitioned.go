package readerpartitioned

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func CreateKafkaReader(readerConfig *kafka.ReaderConfig) *kafka.Reader {
	return kafka.NewReader(*readerConfig)
}

func ReadMessages(reader *kafka.Reader, ctx context.Context, timeout time.Duration) (msgs []kafka.Message) {
	ctx, cancel := context.WithTimeout(ctx, timeout)

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("error while reading message: %v\n", err)
			}
			log.Printf("topic: %s, partition: %d, offset: %d, message: %s\n", m.Topic, m.Partition, m.Offset, string(m.Value))
			msgs = append(msgs, m)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			cancel()
			return msgs
		default:
			continue
		}
	}

}
