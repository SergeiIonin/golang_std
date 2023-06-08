package main

import "fmt"

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
// https://leetcode.com/problems/missing-number/

func main() {
	in0 := []int{3, 0, 1}
	in1 := []int{0, 1}

	r0 := missingNumber(in0)
	r1 := missingNumber(in1)

	fmt.Println(r0)
	fmt.Println(r1)
}

// RT 10 ms (92.72%), Mem 6.2 MB (99.58%) however results seems to be flacky!
func missingNumber(nums []int) int {
	size := len(nums)
	sum := (1 + size) * size / 2
	i := 0
	for i < size/2 {
		sum -= (nums[i] + nums[size-1-i])
		i++
	}
	return sum - nums[i]*(size%2)
}

// func missingNumber(nums []int) int {
//     if len(nums) == 0 {
//         return 0
//     }
//     sum := (1+len(nums))*len(nums)/2
//     for i := range nums {
//         sum-=nums[i]
//     }
//     return sum

// }

// this solution from leetcode performs at 4 ms

// func missingNumber(nums []int) int {
//     n := len(nums)
//     sumAll := (n * (n + 1)) / 2

//     for i := n - 1; i >= 0; i-- {
//         sumAll -= nums[i]
//     }

//     return sumAll
// }
