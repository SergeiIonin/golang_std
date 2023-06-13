package main

import (
	"fmt"
	"sort"
	"time"
)

// https://leetcode.com/problems/3sum/

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

func main() {
	input0 := []int{-1, 0, 1, 2, -1, -4} // [[-1,-1,2],[-1,0,1]]
	input1 := []int{0, 0, 0, 0}          // [[0,0,0]]
	input2 := []int{1, -1, -1, 0}        // [[-1 0 1]]
	input3 := []int{-1, 0, 1, 0}         // [[-1,0,1]]
	input4 := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	input5 := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	input6 := []int{-2, 0, 0, 2, 2}

	start := time.Now()

	res0 := threeSum(input0)
	res4 := threeSum(input4)
	res3 := threeSum(input3)
	res1 := threeSum(input1)
	res2 := threeSum(input2)
	res5 := threeSum(input5)
	res6 := threeSum(input6)

	timeElapsed := time.Since(start)

	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Println(res6)
	fmt.Printf("Time elapsed = %d \n", timeElapsed.Microseconds())

}

func threeSum_(nums []int) [][]int {
	size := len(nums)
	if size < 3 {
		return [][]int{}
	}
	res := make([][]int, 0)

	sort.Ints(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, size-1
		for j < k { // NB if j == len(nums) we don't enter here!
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {

	sort.Ints(nums)

	res := make([][]int, 0, len(nums))

	i := 0
	j := i + 1
	k := len(nums) - 1

	for nums[i] <= 0 && i < k {
		if nums[i]+nums[j] > 0 || j >= k {
			k = len(nums) - 1
			for i < k {
				i++
				if nums[i] != nums[i-1] {
					break
				}
			}
			j = i + 1
			continue
		}
		sum := nums[i] + nums[j] + nums[k]
		if sum < 0 {
			j++
		} else if sum > 0 {
			k--
		} else {
			res = append(res, []int{nums[i], nums[j], nums[k]})
			for j < k {
				j++
				k--
				if nums[j] != nums[j-1] || nums[k] != nums[k+1] {
					break
				}
			}
		}
	}

	return res
}
