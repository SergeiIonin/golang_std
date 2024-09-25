package main

import (
	"fmt"
)

func main() {
	input0 := []int{27,23,90,87,47,75,66,55,58,31,39,71,65,48,4,13,15,1,2,9,68,52,74,51,27,18,29,55,52,34,1,12,78,21,41,65,36,41,72,4,83,18,83,62,67,32,74,57,22,14,27,53,76,18,1,15,52,14,57,18,21,73,73,40,81,54,23,12,16,12,48,24,60,14,43,74,90,11,32,1,74,9,81,34,12,85,49,84,76,44,13,17,86,75,88,87,41,1,54,78,22,83,72,24,72,4,1,86,78,77,20,90,3,53,34,19,46,55,87,78,79,60,60,33,30,65,64,38,83,12,78,36,32,17,81,59,58,24,87,37,63,76,65,56,55,37,70,73,59,37,89,8,43,72,58,66,25,47,2,32,43,60,51,83,72,39,77,29,48,8,64,40,10,66,4,21,51,15,90,83,7,90,20,68,11,24,83,1,69,20,51,87,44,81,21,50,1,5,13,84,20,67,18,24,23,4,8,29,25,78,38,52,30,70,2,77,68,20,56,67,90,15,26,56,32,39,12,19,29,61,20,22,2,41,34,54,37,3,8,34,30,85,55,15,71,20,14,55,86,57,71,90,37,65,49,21,13,14,7,32,34,5,1,55,1,53,48,63,75,74,13,55,52,18,10,46,72,32,19,21,7,88,44,71,8,39,73,43,86,79,56,78,70,89,44,69,76,25,6,83,54,36,19,55,57,57,74,54,33,29,32,10,6,37,39,33,31,48,19,13,45,69,51,23,59,89,30,68,34,15,38,69,64,3,32,52,17,77,70,39,71,61,34,46,2,85,26,73,73,29,35,11,35,27,23,75,72,9,61,78,78,73,11,49,8,29,2,30,25,39,19,22,12,51,17,26,64,5,22,12,21,84,53,1,3,53,62,70,23,34,75,5,31,89,22,7,15,36,32,68,3,27,56,81,65,52,44,34,62,11,69,17,14,60,19,63,44,12,18,46,59,46,60,11,12,10,2,65,58,44,58,68,70,85,86,78,78,50,6,73,53,71,33,30,14,60,12,14,48,6,72,81,39}
	res := maxOperations(input0, 68) // 126
	fmt.Printf("res = %d", res)
}

/*
https://leetcode.com/problems/max-number-of-k-sum-pairs
You are given an integer array nums and an integer k.
In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
Return the maximum number of operations you can perform on the array.

Example 1:

Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.

Example 2:

Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.

Constraints:
    1 <= nums.length <= 10^5
    1 <= nums[i] <= 10^9
    1 <= k <= 10^9
*/
func maxOperations(nums []int, k int) int {
	dict := make(map[int]int)
	l := 0
	r := len(nums)-1
	res := 0

	min := func(l, r int) int {
		if l < r {
			return l
		}
		return r
	}
	addOccurrenceOrDelete := func(dict map[int]int, key, occ int) {
		if occ == 0 {
			delete(dict, key)
		} else {
			dict[key] = occ
		}
	}

	for ; l < r; {
		if nums[l] >= k {
			l++
			continue
		} else if nums[r] >= k {
			r--
			continue
		}
		wanted := k - nums[l]
		if nums[r] == wanted {
			l++
			r--
			res++
		} else {
			occs, ok := dict[wanted]
			if ok {
				addOccurrenceOrDelete(dict, wanted, occs-1)
				l++
				res++
			}
			dict[nums[r]]++
			r--
		}
	}
	if l == r {
		dict[nums[r]]++
	}
	for elem, occs := range dict {
		wanted := k - elem
		occsWanted, ok := dict[wanted]
		if ok {
			var numMatches int
			if elem == wanted {
				numMatches = occsWanted / 2
				addOccurrenceOrDelete(dict, elem, 0)
			} else {
				numMatches = min(occs, occsWanted)
				addOccurrenceOrDelete(dict, elem, occs-numMatches)
				addOccurrenceOrDelete(dict, wanted, occsWanted-numMatches)
			}
			res += numMatches
		}
	}
	return res
}
