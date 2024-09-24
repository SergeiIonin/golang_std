package main

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
    tests := []struct {
        name     string
        nums     []int
        expected []int
    }{
        {"Test 1", []int{0,1,0,3,12}, []int{1,3,12,0,0}},
        {"Test 2", []int{0}, []int{0}},
        {"Test 3", []int{1,0,1}, []int{1,1,0}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            moveZeroes(tt.nums)
            if !slices.Equal(tt.nums, tt.expected) {
                t.Errorf("got %v, want %v", tt.nums, tt.expected)
            }
        })
    }
}