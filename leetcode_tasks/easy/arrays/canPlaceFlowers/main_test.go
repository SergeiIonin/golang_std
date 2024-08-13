package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
    tests := []struct {
        name     string
        flowers  []int
        n        int
        expected bool
    }{
        {"Test 1", []int{1, 0, 0, 0, 1}, 1, true},
        {"Test 2", []int{1, 0, 0, 0, 1}, 2, false},
		{"Test 3", []int{1,0,0,0,1}, 1, true},
		{"Test 4", []int{1}, 0, true},
		{"Test 5", []int{1}, 1, false},
		{"Test 6", []int{1,0,0,0,1,0,0}, 2, true},
		{"Test 7", []int{0,0,1,0,0}, 1, true},
		{"Test 8", []int{0,1,0,1,0,1,0,0}, 1, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := canPlaceFlowers(tt.flowers, tt.n)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}