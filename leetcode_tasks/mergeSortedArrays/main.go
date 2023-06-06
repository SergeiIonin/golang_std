package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	r := merge(nums1, m, nums2, n)

	fmt.Println(r)
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		return nums2
	}
	if n == 0 {
		return nums1[:m]
	}
	{
		head1 := nums1[0]
		head2 := nums2[0]
		if head1 < head2 {
			nums1 = append([]int{head1}, merge(nums1[1:], m-1, nums2, n)...)
			return nums1
		} else {
			nums2 = append([]int{head2}, merge(nums1, m, nums2[1:], n-1)...)
			return nums2
		}
	}
}
