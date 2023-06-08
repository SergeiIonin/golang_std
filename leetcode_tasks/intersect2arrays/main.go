package main

import "fmt"

// https://leetcode.com/problems/intersection-of-two-arrays-ii

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must appear as many times as it shows in both arrays
// and you may return the result in any order.

func main() {
	in1_0 := []int{1, 2, 2, 1}
	in2_0 := []int{2, 2}

	in1_1 := []int{4, 9, 5}
	in2_1 := []int{9, 4, 9, 8, 4}

	res0 := intersect(in1_0, in2_0)
	res1 := intersect(in1_1, in2_1)

	fmt.Println(res0)
	fmt.Println(res1)
}

// pretty similar to the best solution on leetcode
func intersect(nums1 []int, nums2 []int) []int {

	size := len(nums1)
	if len(nums2) < len(nums1) {
		size = len(nums2)
	}
	var res []int

	hash := make(map[int]int, size)

	if len(nums1) < len(nums2) {
		addToMap(hash, nums1)
		res = fillResArray(hash, nums2)
	} else {
		addToMap(hash, nums2)
		res = fillResArray(hash, nums1)
	}

	return res

}

func addToMap(h map[int]int, nums []int) {
	for i := 0; i < len(nums); i++ {
		h[nums[i]]++
	}
}

func fillResArray(h map[int]int, nums []int) []int {
	res := make([]int, 0, len(h))
	for j := 0; j < len(nums); j++ {
		v, ok := h[nums[j]]
		if ok {
			if v >= 1 {
				h[nums[j]]--
				res = append(res, nums[j])
			}
		}
	}
	return res
}
