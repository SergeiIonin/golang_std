package main

import (
    "testing"
)

func TestGcdOfStrings(t *testing.T) {
    tests := []struct {
        name     string
        str1     string   
        str2     string
        expected string
    }{
        {"Test 1", "ABCABC", "ABC", "ABC"},
        {"Test 2", "ABABAB", "ABAB", "AB"},
        {"Test 3", "LEET", "CODE", ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := gcdOfStrings(tt.str1, tt.str2)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}