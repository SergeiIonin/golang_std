package main

import (
    "testing"
)

func TestReverseWords(t *testing.T) {
    tests := []struct {
        name     string
        word     string
        expected string
    }{
        {"Test 1", "the sky is blue", "blue is sky the"},
        {"Test 2", "  hello world  ", "world hello"},
		{"Test 3", "a good   example", "example good a"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := reverseWords(tt.word)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}