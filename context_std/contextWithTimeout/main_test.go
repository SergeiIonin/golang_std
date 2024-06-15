package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func Test_getResultOrTimeout(t *testing.T) {
	type args struct {
		ctx     context.Context
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"timeout elapsed first",
			args{
				context.Background(),
				3 * time.Second,
			},
			"",
			true,
		},
		{
			"result returned first",
			args{
				context.Background(),
				6 * time.Second,
			},
			"todo bien, ¿y tu?",
			false,
		},
	}
	for _, tt := range tests {
		log.Println("Running tests...")
		t.Run(tt.name, func(t *testing.T) {
			got, err := getResultOrTimeout(tt.args.ctx, tt.args.timeout)
			if (err != nil) && tt.wantErr {
				log.Printf("getResultOrTimeout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getResultOrTimeout() = %v, want %v", got, tt.want)
			}
		})
	}
}
