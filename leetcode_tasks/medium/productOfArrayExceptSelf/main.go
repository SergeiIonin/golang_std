package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

func main() {
	in0 := []int{1, 2, 3, 4}
	r0 := productExceptSelf(in0) // [24, 12, 8, 6]
	fmt.Println(r0)
}

func productExceptSelf(nums []int) []int {

	res := make([]int, len(nums))
	multiplyer := 1

	// we first calculate product of all elems before the current, hence the last product is ready
	for i := 0; i < len(nums); i++ {
		res[i] = multiplyer
		multiplyer = multiplyer * nums[i]
	}

	// then we multiply the intermediate results by the growing mulptiplyer, going from the last elem to head
	multiplyer = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * multiplyer
		multiplyer = multiplyer * nums[i]
	}

	return res
}
