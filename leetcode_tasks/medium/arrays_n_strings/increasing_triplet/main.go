package main

import "fmt"

func main() {
	in := []int{0,4,2,1,0,-1,-3}
	res := increasingTriplet(in)
	fmt.Printf("increasingTriplet found: %v", res)
}

// https://leetcode.com/problems/increasing-triplet-subsequence
/* Constraints:
    1 <= nums.length <= 5 * 10^5
    -2^31 <= nums[i] <= 2^31 - 1
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
    res := make([]int, 0, 3)
	minInd := 0
	res = append(res, minInd)

	for i, elem := range nums {
		if elem < nums[minInd] {
			minInd = i
		}
		if len(res) == 1 {
			res[0] = minInd
			if elem > nums[minInd] {
				res = append(res, i)
			}
		} else if len(res) == 2 {
			if elem > nums[res[1]] {
				return true
			} else if elem > nums[minInd] && elem < nums[res[1]] {
				res[0] = minInd
				res[1] = i
			}
		}
	}
	
	return false
}
