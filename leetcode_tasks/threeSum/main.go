package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	input0 := []int{-1, 0, 1, 2, -1, -4} // [[-1,-1,2],[-1,0,1]]
	input1 := []int{0, 0, 0, 0}          // [[0,0,0]]
	input2 := []int{1, -1, -1, 0}        // [[-1 0 1]]
	input3 := []int{-1, 0, 1, 0}         // [[-1,0,1]]
	input4 := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	input5 := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}

	//[-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]

	start := time.Now()

	res0 := threeSum(input0)
	res4 := threeSum(input4)
	res3 := threeSum(input3)
	res1 := threeSum(input1)
	res2 := threeSum(input2)
	res5 := threeSum(input5)

	timeElapsed := time.Since(start)

	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)
	fmt.Printf("Time elapsed = %d \n", timeElapsed.Microseconds())

}

func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return [][]int{}
	}

	res := make([][]int, 0, len(nums))

	sort.Ints(nums)

	i := 0
	j := 1

	for i < len(nums)-2 {

		zero := nums[i]
		if zero > 0 {
			break
		}
		one := nums[j]

		if j == len(nums)-1 {
			if i >= len(nums)-3 {
				break
			}
			i++
			nums = nums[:len(nums)-1]
			j = i + 1
			continue
		}

		if i > 0 {
			if zero == nums[i-1] { // 'zero' elem can't repeat
				i++
				j = j + 1
				continue
			}
		}

		if j != i+1 && one == nums[j-1] { // 'one' never equals to previous elem except 'zero'
			j = j + 1
			continue
		}

		two := -zero - one

		if two < 0 {
			i++
			j = i + 1
			continue
		}
		for k := j + 1; k < len(nums); k++ {
			elem := nums[k]
			if elem > two {
				break
			}
			if elem == two {
				triplet := []int{zero, one, two}
				res = append(res, triplet)
				break
			}
		}
		j++

	}

	return res
}
