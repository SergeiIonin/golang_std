package dial

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"slices"
	"testing"
)

func TestKafkaKraftOnPremise_test(t *testing.T) {

	// Connect to Kafka
	kafkaAddr := "localhost:9092"
	t.Logf("Kafka address: %s", kafkaAddr)

	client := &kafka.Client{
		Addr: kafka.TCP(kafkaAddr),
		//Timeout:   30 * time.Second,
		Transport: nil,
	}

	_, err := client.Heartbeat(context.Background(), &kafka.HeartbeatRequest{
		Addr:            kafka.TCP(kafkaAddr),
		GroupID:         "",
		GenerationID:    0,
		MemberID:        "",
		GroupInstanceID: "",
	})
	if err != nil {
		log.Fatal("failed to heartbeat:", err)
	}
	log.Println("Heartbeat is successful")

	/*log.Println("WAITING...")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(600 * time.Second)
		wg.Done()
	}()
	wg.Wait()*/

	topics := []string{"test_1", "test_2", "test_3"}
	topicsConfigs := make([]kafka.TopicConfig, 0, len(topics))
	for _, topic := range topics {
		conf := kafka.TopicConfig{
			topic,
			1,
			1,
			nil,
			nil,
		}
		topicsConfigs = append(topicsConfigs, conf)
	}
	createTopicsReq := &kafka.CreateTopicsRequest{
		Addr:         kafka.TCP(kafkaAddr),
		Topics:       topicsConfigs,
		ValidateOnly: false,
	}

	_, err = client.CreateTopics(context.Background(), createTopicsReq)
	if err != nil {
		log.Fatal("failed to create topics:", err)
	}

	readerConfig := kafka.ReaderConfig{
		Brokers:  []string{kafkaAddr},
		Topic:    "test_1",
		GroupID:  "test_twks",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	reader := kafka.NewReader(readerConfig)

	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaAddr),
		Topic:    "test_1",
		Balancer: &kafka.LeastBytes{},
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("AAA"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("BBB"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("CCC"),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	msgs := make([]string, 9)

	for _, _ = range msgs {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		msgs = append(msgs, string(m.Value))
		log.Printf("message: %v", m)
	}

	log.Printf("messages read: %v", msgs)

	res := slices.Equal([]string{"AAA", "BBB", "CCC"}, msgs)

	if !res {
		log.Fatal("assertion for msgs failed")
	}

	/*wg := &sync.WaitGroup{}
	wg.Add(1)
	t_start := time.Now()
	go func() {
		time.Sleep(600 * time.Second)
		wg.Done()
	}()
	wg.Wait()
	t_end := time.Now()
	defer func() {
		timeElapsed := t_end.Sub(t_start)
		log.Printf("Time elapsed: %s", timeElapsed)
		if err := kafkaC.Terminate(ctx); err != nil {
			log.Fatalf("Could not stop kafka: %s", err)
		}
	}()*/

}
