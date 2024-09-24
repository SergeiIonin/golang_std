package main

import (
    "testing"
)

func TestReverseWords(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        expected bool
    }{
        {"Test 1", []int{1,2,3,4,5}, true},
        {"Test 2", []int{5,4,3,2,1}, false},
        {"Test 3", []int{2,1,5,0,4,6}, true},
        {"Test 4", []int{0,4,2,1,0,-1,-3}, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := increasingTriplet(tt.nums)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}