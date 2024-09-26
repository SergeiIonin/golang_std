package main

import (
	"testing"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k 		 int
		expected int
	}{
		{"Test 1", []int{1,1,1,0,0,0,1,1,1,1,0}, 2, 6},
		{"Test 2", []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3, 10},
		{"Test 3", []int{0,0,1,1,1,0,0}, 0, 3},
		{"Test 4", []int{0,0,0,0}, 0, 0},
		{"Test 5", []int{1,1,1,0,0,0,1,1,1,1}, 0, 4},
		{"Test 6", []int{0,1,1}, 0, 2},
		{"Test 7", []int{0}, 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := longestOnes(tt.nums, tt.k)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}