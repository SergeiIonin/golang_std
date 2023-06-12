package main

import (
	"fmt"
)

func main() {
	in0 := []int{100, 4, 200, 1, 3, 2}
	res0 := longestConsecutive(in0)
	fmt.Println(res0)

	in1 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res1 := longestConsecutive(in1)
	fmt.Println(res1)
}

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	max := nums[0]

	hashmap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]]++
		if max < nums[i] {
			max = nums[i]
		}
	}

	maxLength := 0
	length := 0

	for i := 0; i < len(nums); i++ {
		var elem int

		if max-nums[i] < maxLength {
			continue
		}

		value := hashmap[nums[i]]
		_, hasLeastUpperBound := hashmap[nums[i]+maxLength]
		if value > 0 && hasLeastUpperBound {
			elem = nums[i]
		} else {
			continue
		}

		for {
			_, ok := hashmap[elem]
			if ok {
				hashmap[elem]--
				length++
				elem++
			} else {
				break
			}
		}

		if length > maxLength {
			maxLength = length
		}
		length = 0
	}

	if length > maxLength {
		maxLength = length
	}

	return maxLength
}
