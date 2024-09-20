package main

import (
    "testing"
    "slices"
)

func TestKidsWithCandies(t *testing.T) {
    tests := []struct {
        name     string
        candies  []int
        n        int
        expected []bool
    }{
        {"Test 1", []int{12,1,12}, 10, []bool{true, false, true}},
        {"Test 2", []int{4,2,1,1,2}, 1, []bool{true,false,false,false,false}},
		{"Test 3", []int{2,3,5,1,3}, 3, []bool{true,true,true,false,true}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := kidsWithCandies(tt.candies, tt.n)
            if !slices.Equal(result, tt.expected) {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}