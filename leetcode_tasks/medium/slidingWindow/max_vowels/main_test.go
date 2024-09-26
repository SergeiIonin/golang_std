package main

import (
	"testing"
)

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		k 		 int
		expected int
	}{
		{"Test 1", "abciiidef", 3, 3},
		{"Test 2", "aeiou", 2, 2},
		{"Test 3", "leetcode", 3, 2},
		{"Test 4", "twfsikrausdeuelcgiupwktszz", 5, 3},
		{"Test 5", "weallloveyou", 7, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxVowels(tt.str, tt.k)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}