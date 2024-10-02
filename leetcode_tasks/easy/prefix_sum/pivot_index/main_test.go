package main

import "testing"

func TestPivotIndex(t *testing.T) {
    tests := []struct {
        name     string
        in      []int
        expected int
    }{
        {"Test 1", []int{1,7,3,6,5,6}, 3},
        {"Test 2", []int{1,2,3}, -1},
        {"Test 3", []int{2,1,-1}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := pivotIndex(tt.in)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}