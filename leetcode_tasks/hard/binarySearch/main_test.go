package main

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {

	inputs := []TestInputAndResult{
		{
			[]int{1, 3},
			[]int{2, 7},
			2.5,
			"two arrays of size 2",
		},
	}
	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			got := findMedianSortedArrays(input.nums1, input.nums2)
			if got != input.res {
				fmt.Println("test ", input.name)
				fmt.Println("got = ", got)
				fmt.Println("expected = ", input.res)
				t.Error("test failed")
			}
		})
	}

}

type TestInputAndResult struct {
	nums1 []int
	nums2 []int
	res   float64
	name  string
}
