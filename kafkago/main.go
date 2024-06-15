package main

import (
	"context"
	"log"
	"time"

	"github.com/SergeiIonin/golang_std/kafkago/dial"
)

func main() {
	context := context.Background()
	timeout := 3 * time.Second
	_, err := dial.DialKafka(context, timeout, "tcp", "127.0.0.1:9092")
	if err != nil {
		log.Printf("error occured is %s", err)
		return
	}
	log.Println("todo bien")
}
