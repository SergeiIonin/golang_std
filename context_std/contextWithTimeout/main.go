package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("in case of timeout elapsed first, context will be canceled")
	result, err := getResultOrTimeout(ctx, 3*time.Second)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}
	ctx = context.Background()
	fmt.Println("in case of result returned first")
	result, err = getResultOrTimeout(ctx, 6*time.Second)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}
}

func getResultOrTimeout(ctx context.Context, timeout time.Duration) (string, error) {
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(ctx, timeout)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "todo bien, ¿y tu?"
	}()

	select {
	case <-ctx.Done():
		cancel()
		return "", errors.New("timeout")
	case result := <-ch:
		return result, nil
	}
}
