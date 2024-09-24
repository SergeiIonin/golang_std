package main

import "fmt"

func main() {
	in := []int{0, 1, 0, 3, 12}
	moveZeroes(in)
	fmt.Printf("res = %v\n", in)
}

/*
https://leetcode.com/problems/move-zeroes/
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]

Constraints:

	1 <= nums.length <= 10^4
	-2^31 <= nums[i] <= 2^31 - 1
*/
func moveZeroes(nums []int) {
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] == 0 && nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		if nums[i] != 0 {
			i++
		}
		j++
	}
}
