package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.

func main() {
	in0 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	r0 := removeDuplicates(in0)
	fmt.Println(r0)
	fmt.Println(in0)

	in1 := []int{1, 1}
	r1 := removeDuplicates(in1)
	fmt.Println(r1)
	fmt.Println(in1)

	in2 := []int{1, 1, 1}
	r2 := removeDuplicates(in2)
	fmt.Println(r2)
	fmt.Println(in2)
}

func removeDuplicates(nums []int) int {
	start := -1
	end := -1
	i := 0
	j := 1
	lenDupls := 0
	for i < len(nums) && j < len(nums) {
		if nums[j] < nums[i] {
			break
		}
		if nums[i] == nums[j] {
			if start == -1 {
				start = j
			}
			lenDupls++ // it was a caveat! we shouldn't incr lenDupls in the 'if' below
		}
		if start != -1 && nums[i] != nums[j] {
			end = j
			dupls := append([]int{}, nums[start:end]...)
			nums = append(nums[:start], nums[end:]...)
			nums = append(nums, dupls...)
			i = start
			j = i + 1
			start = -1
			continue
		}
		i++
		j++
	}
	return len(nums) - lenDupls
}

// this is the most performant solution on leetcode
func removeDuplicates_(nums []int) int {
	curr := nums[0]
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != curr {
			curr = nums[i]
			nums[j] = curr
			j++
		}
	}
	return j
}
