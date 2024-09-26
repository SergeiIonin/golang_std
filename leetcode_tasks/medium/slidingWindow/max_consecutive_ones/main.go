package main

import (
	"fmt"
)

func main() {
	in := []int{0,1,1}
	res := longestOnes(in, 0)
	fmt.Printf("res = %d", res)
}
// https://leetcode.com/problems/max-consecutive-ones-iii
/* Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Constraints:
    1 <= nums.length <= 10^5
    nums[i] is either 0 or 1.
    0 <= k <= nums.length
 */

 func longestOnes(nums []int, k int) int {
	if len(nums) == 1 {
		if nums[0] == 0 && k > 0 {
			return 1
		} else {
			return 1
		}
	}

	fmt.Printf("nums = %v \n", nums)

	l := 0
	r := 1
	count := 0
	max := count
	quotes := k

	moveWindow := func() {
		for ; l < len(nums) ; l++ {
			if count > 0 {
				count--
			}
			if nums[l] == 0 {
					l++
					if k != 0 {
						quotes++
					} else {
						r++
					}
					return
			}
		}
	}

	if quotes != 0 {
		if nums[0] == 0 {
			quotes--
		}
		count++
		max = count
	} else {
		for j, elem := range nums {
			if elem == 1 {
				l = j
				count++
				max = count
				r = l+1
				break
			}
		}
	}

	for ; r < len(nums); {
		fmt.Printf("l = %d, r = %d, window = %v, count = %d, quotes = %d \n", l, r, nums[l:r], count, quotes)
		fmt.Printf("nums[%d] = %d \n", r, nums[r])
		if nums[r] == 1 {
			count++
			r++
		} else if quotes > 0 {
			quotes--
			count++
			r++
		} else if quotes == 0 {
			fmt.Printf("quotes are exhausted, r = %d \n", r)
			if count > max {
				max = count
				fmt.Printf("max = %d \n", max)
			}
			moveWindow()
			fmt.Printf("[after move] l = %d, r = %d, count = %d, quotes = %d \n", l, r, count, quotes)
		}
	}

	if count > max {
		max = count
	}

	return max
 }
 