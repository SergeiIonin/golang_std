package main

import (
	"fmt"
)

func main() {
	in0 := []int{100, 4, 200, 1, 3, 2}
	res0 := longestConsecutive(in0)
	fmt.Println(res0) // 4

	in1 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	res1 := longestConsecutive(in1)
	fmt.Println(res1) // 9

	in2 := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6} // 3..9
	res2 := longestConsecutive(in2)
	fmt.Println(res2) // 7
}

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	min := nums[0]
	indexMin := 0

	hashmap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]]++
		if max < nums[i] {
			max = nums[i]
			continue
		}
		if min > nums[i] {
			min = nums[i]
			indexMin = i
		}
	}

	maxLength := 1
	length := 0

	i := indexMin

	var elem int
	maxIndex := len(nums) - 1
	count := 2

	for count != 0 {
		if i == indexMin {
			count--
		}

		num := nums[i]

		_, ok := hashmap[num]
		next := num + maxLength
		if next > max {
			i = getIndex(maxIndex, i+1)
			continue
		}
		_, hasLeastUpperBound := hashmap[next]
		if ok && hasLeastUpperBound {
			elem = num

			for {
				_, ok := hashmap[elem]
				if ok {
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
		i = getIndex(maxIndex, i+1)
	}

	if length > maxLength {
		maxLength = length
	}

	return maxLength
}

func getIndex(maxIndex, currentIndex int) int {
	if currentIndex <= maxIndex {
		return currentIndex
	} else {
		return currentIndex - maxIndex - 1
	}
}
