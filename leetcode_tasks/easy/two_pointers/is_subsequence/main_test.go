package main

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s     	 string
		t 		 string
		expected bool
	}{
		{"Test 1", "abc", "ahbgdc", true},
		{"Test 2", "axc", "ahbgdc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubsequence(tt.s, tt.t)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
