package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-the-duplicate-number

// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

// There is only one repeated number in nums, return this repeated number.

// You must solve the problem without modifying the array nums and uses only constant extra space.

func main() {

	in0 := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(in0)) // 2

	in1 := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(in1)) // 3

}

// here's the impl of Hare and Tortoise algorithm
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = 0

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
