package dial

import (
	"context"
	"errors"
	"time"

	"github.com/segmentio/kafka-go"
)

func DialKafka(ctx context.Context, timeout time.Duration, network string, broker string) (*kafka.Conn, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	cchan := make(chan *kafka.Conn)
	errChan := make(chan error)

	// return error if kafka.Dial takes more than 30 seconds

	go func(ch chan *kafka.Conn, echan chan error) {
		conn, err := kafka.Dial(network, broker)
		if err != nil {
			errChan <- err
		}
		ch <- conn
	}(cchan, errChan)

	select {
	case <-ctx.Done():
		cancel()
		return nil, errors.New("timeout dialiig Kafka elapsed!")
	case conn := <-cchan:
		return conn, nil
	case err := <-errChan:
		return nil, err
	}
}
