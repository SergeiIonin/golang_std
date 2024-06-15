package writerpartitioned

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateKafkaWriter(writerConfig *kafka.WriterConfig) *kafka.Writer {
	return kafka.NewWriter(*writerConfig)
}

func WriteMessages(writer *kafka.Writer, messages []kafka.Message) error {
	for _, message := range messages {
		err := writer.WriteMessages(context.Background(), message)
		if err != nil {
			log.Printf("could not write message %v", err)
			return err
		}
	}
	return nil
}

func GenereateMessagesForPartitionedTopic(messages []string) []kafka.Message {
	var kafkaMessages []kafka.Message
	for ind, message := range messages {
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Key:   []byte(string(ind)),
			Value: []byte(message),
		})
	}
	return kafkaMessages
}
