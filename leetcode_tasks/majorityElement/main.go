package main

import "fmt"

func main() {
	in0 := []int{2, 2, 1, 1, 1, 2, 2}
	r := majorityElement(in0)
	fmt.Println(r)
}

// RT 18 ms (63.66%), Mem 5.8 MB (99.93%)
func majorityElement(nums []int) int {
	hash := make(map[int]int, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}
	threshold := len(nums) / 2
	res := 0
	for k, v := range hash {
		if v > threshold {
			return k
		}
	}
	return res
}

// solution from leetcode:
// func majorityElement_(nums []int) int {
// 	candidate := nums[0]
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		if count == 0 {
// 			candidate = nums[i]
// 			count++
// 		} else if candidate == nums[i] {
// 			count++
// 		} else {
// 			count--
// 		}
// 	}

// 	return candidate
// }
