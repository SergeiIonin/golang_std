package context_with_cancellation

import (
	"context"
	"testing"
)

func Test_CtxWithCancelationAndDefault(t *testing.T) {

	t.Run("we can use select with default and read ctx.Done channel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		counter := 0

		incr := func(counter *int, ctx context.Context) {
			*counter++
			if *counter == 5 {
				cancel()
			}
		}

		for {
			select {
			case <-ctx.Done():
				t.Log("ctx.Done")
				return
			default:
				t.Log("default")
				incr(&counter, ctx)
				continue
			}
		}

		if counter != 5 {
			t.Errorf("counter should be 5, but got %d", counter)
		}

	})
}
