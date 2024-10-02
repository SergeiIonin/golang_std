package main

import "testing"

func TestLargestAltitude(t *testing.T) {
    tests := []struct {
        name     string
        in      []int
        expected int
    }{
        {"Test 1", []int{-5,1,5,0,-7}, 1},
        {"Test 2", []int{-4,-3,-2,-1,4,3,2}, 0},
        {"Test 3", []int{52,-91,72}, 52},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := largestAltitude(tt.in)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}