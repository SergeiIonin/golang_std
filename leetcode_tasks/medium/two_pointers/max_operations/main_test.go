package main

import (
	"testing"
)

func TestMaxOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k 		 int
		expected int
	}{
		{"Test 1", []int{2,2,2,3,1,1,4,1}, 4, 2},
		{"Test 2", []int{4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4}, 2, 2},
		{"Test 3", []int{3,1,3,4,3}, 6, 1},
		{"Test 4", []int{13,16,49,4,56,64,83,35,20,73,53,67,73,73,17,28,20,16,55,16,20,66,13,46,9,14,52,70,46,66,40,21,5,88,48,21,21,44,27,56,75,58,57,15,27,4,51,77,17,21,65,17,62,84,71,78,10,67,49,8,47,55,41,86,43,48,69,58,62,27,38,24,12,82,38,62,82,32,29,27,38,37,78,9,74,90,64,16,37,22,37,46,20,47,31,16,81,28,82,39,86,59,11,78,12,13,71,49,69,1,37,24,79,32,25,67,42,30,16,23,51,66,72,20,11,90,34,81,10,86,51,68,10,62,59,33,49,30,81,69,80,79,54,78,87,44,40,47,78,44,30,23,41,89,35,6,88,79,14,10,27,54,83,36,78,82,51,1,48,28,72,34,41,32,47,32,42,3,25,78,28,37,27,77,32,83,29,56,86,80,50,59,44,51,25,41,18,83,62,4,16,80,72,7,34,21,81,15,20,35,15,46,55,81,2,36,70,87,52,19,76,18,27,81,19,78,36,6,84,32,27,7,70,67,87,90,37,75,80,72,60,68,6,72,12,90,83,20,42,36,62,45,49,45,56,51,66,48,30,49,58,9,4,56,53,30,22,7,43,23,89,46,81,61,37,78,30,9,55,43,76,65,68,64,31,1,80,39,72,45,37,88,54,23,89,13,68,26,75,86,82,69,15,25,57,38,89,70,47,4,7,11,57,64,10,73,15,16,58,5,39,60,86,50,1,85,7,40,37,58,57,52,2,13,14,73,83,29,35,76,24,55,52,75,74,38,59,73,36,90,66,61,74,74}, 77, 107},
		{"Test 5", []int{82,85,16,61,14,3,81,59,31,68,23,65,70,90,67,49,41,31,59,74,62,32,12,16,23,38,46,70,66,9,85,75,9,43,10,71,30,9,26,9,49,83,82,55,38,59,88,12,15,71,6,62,53,88,88,33,28,34,44,80,86,24,67,20,80,81,71,6,66,48,53,7,22,9,1,56,78,14,45,80,62,3,72,73,13,71,23,26,71,63,39,29,88,36,27,70,21,6,43,7,82,62,21,24,37,89,53,3,66,21,47,7,38,12,21,31,16,44,45,87,37,5,11,25,77,16,78,81,80,86,12,70,29,63,79,73,50,74,27,50,64,3,53,78,69,77,55,77,59,4,70,32,56,54,89,56,53,60,33,25,41,30,9,51,3,22,19,51,58,17,9,60,81,12,16,82,65,54,17,89,15,66,15,17,57,39,22,80,10,86,85,65,14,64,87,64,54,13,5,18,30,26,34,3,30,42,4,51,41,78,49,38,17,85,17,59,51,46,86,59,53,4,82,77,49,22,6,20,36,70,13,24,61,48,79,14,19,35,84,76,87,46,30,6,45,86,71,73,61,38,62,27,90,85,36,62,58,14,6,52,45,82,46,85,89,69,87,62,86,52,51,25,62,55,86,18,32,37,26,56,69,87,29,68,13,17,22,47,16,33,65,24,82,1,34,53,57,27,67,40,25}, 80, 91},
		{"Test 6", []int{27,23,90,87,47,75,66,55,58,31,39,71,65,48,4,13,15,1,2,9,68,52,74,51,27,18,29,55,52,34,1,12,78,21,41,65,36,41,72,4,83,18,83,62,67,32,74,57,22,14,27,53,76,18,1,15,52,14,57,18,21,73,73,40,81,54,23,12,16,12,48,24,60,14,43,74,90,11,32,1,74,9,81,34,12,85,49,84,76,44,13,17,86,75,88,87,41,1,54,78,22,83,72,24,72,4,1,86,78,77,20,90,3,53,34,19,46,55,87,78,79,60,60,33,30,65,64,38,83,12,78,36,32,17,81,59,58,24,87,37,63,76,65,56,55,37,70,73,59,37,89,8,43,72,58,66,25,47,2,32,43,60,51,83,72,39,77,29,48,8,64,40,10,66,4,21,51,15,90,83,7,90,20,68,11,24,83,1,69,20,51,87,44,81,21,50,1,5,13,84,20,67,18,24,23,4,8,29,25,78,38,52,30,70,2,77,68,20,56,67,90,15,26,56,32,39,12,19,29,61,20,22,2,41,34,54,37,3,8,34,30,85,55,15,71,20,14,55,86,57,71,90,37,65,49,21,13,14,7,32,34,5,1,55,1,53,48,63,75,74,13,55,52,18,10,46,72,32,19,21,7,88,44,71,8,39,73,43,86,79,56,78,70,89,44,69,76,25,6,83,54,36,19,55,57,57,74,54,33,29,32,10,6,37,39,33,31,48,19,13,45,69,51,23,59,89,30,68,34,15,38,69,64,3,32,52,17,77,70,39,71,61,34,46,2,85,26,73,73,29,35,11,35,27,23,75,72,9,61,78,78,73,11,49,8,29,2,30,25,39,19,22,12,51,17,26,64,5,22,12,21,84,53,1,3,53,62,70,23,34,75,5,31,89,22,7,15,36,32,68,3,27,56,81,65,52,44,34,62,11,69,17,14,60,19,63,44,12,18,46,59,46,60,11,12,10,2,65,58,44,58,68,70,85,86,78,78,50,6,73,53,71,33,30,14,60,12,14,48,6,72,81,39}, 68, 126},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxOperations(tt.nums, tt.k)
			if res != tt.expected {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}