package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{"Test 1", []int{1,8,6,2,5,4,8,3,7}, 49},
		{"Test 2", []int{1,1}, 1},
		{"Test 3", []int{1, 5, 16, 16, 5, 3}, 16},
		{"Test 4", []int{1, 5, 16, 14, 5, 3}, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxArea(tt.height)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
