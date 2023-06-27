package main

import "sort"

// https://leetcode.com/problems/find-the-duplicate-number

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.

// You must solve the problem without modifying the array nums and uses only constant extra space.

func main() {

}

// it works, but it's very slow and even not very good memory-wise
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
