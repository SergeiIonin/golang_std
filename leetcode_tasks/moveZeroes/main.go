package main

import "fmt"

// https://leetcode.com/problems/move-zeroes/

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

func main() {
	in0 := []int{0, 1, 0, 3, 12}
	in1 := []int{0}
	in2 := []int{1, 0}

	moveZeroes(in0)
	moveZeroes(in1)
	moveZeroes(in2)

	fmt.Println(in0)
	fmt.Println(in1)
	fmt.Println(in2)
}

func moveZeroes(nums []int) {
	i := 0
	j := 1
	for i < len(nums) && j < len(nums) {
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

// func moveZeroes(nums []int) {
// 	i := 0
// 	j := 1
// 	for i < len(nums) && j < len(nums) {
// 		if nums[i] == 0 && nums[j] != 0 {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			i++
// 		}
// 		if nums[i] != 0 {
// 			i++
// 		}
// 		j++
// 	}
// }
