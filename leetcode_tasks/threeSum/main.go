package main

import (
	"fmt"
	"sort"
)

func main() {
	input0 := []int{-1, 0, 1, 2, -1, -4} // [[-1,-1,2],[-1,0,1]]
	input1 := []int{0, 0, 0, 0}          // [[0,0,0]]
	input2 := []int{1, -1, -1, 0}        // [[-1 0 1]]
	input3 := []int{-1, 0, 1, 0}         // [[-1,0,1]]
	input4 := []int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}
	input5 := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6} // [[-1,0,1]]

	//[-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]

	res4 := threeSum(input4)
	res3 := threeSum(input3)
	res1 := threeSum(input1)
	res0 := threeSum(input0)
	res2 := threeSum(input2)
	res5 := threeSum(input5)

	fmt.Println(res0)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println(res4)
	fmt.Println(res5)

}

// [-4 -1 -1 0 1 2]

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	res := make([][]int, 0, len(nums))

	nums = filterNums(nums)

	i := 0
	j := 1

	temp := make([]int, 0, 3)

	for i < len(nums)-2 {

		if j == len(nums)-1 {
			if i >= len(nums)-3 {
				break
			}
			i++
			temp = make([]int, 0, 3)
			if nums[i] == nums[i-1] {
				i++
				j = i + 1
				continue
			}
			j = i + 1
		}

		if len(temp) == 0 {
			temp = append(temp, nums[i])
		}

		if len(temp) == 1 {
			one := nums[j]
			cond1 := one == nums[j-1] && i != j-1
			cond2 := !(temp[0] == 0 && one == 0) // temp array is not all zeros
			if cond1 && cond2 {
				j++
				continue
			}
			temp = append(temp, one)
		}

		if len(temp) == 2 {
			target := -temp[0] - temp[1]

			for k := j + 1; k < len(nums); k++ {
				if nums[k] == target {
					triplet := []int{temp[0], temp[1], target}
					res = append(res, triplet)
					break
				}
			}
			temp = make([]int, 0, 3)
			j++
		}
	}

	return res
}

func threeSumOld(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	res := make([][]int, 0, len(nums))

	nums = filterNums(nums)

	i := 0
	j := 1
	k := 2

	temp := make([]int, 0, 3)
	for i < len(nums)-2 {

		if i > 0 && nums[i] == nums[i-1] {
			i++
			j = i + 1
			k = j + 1
		}

		if k == len(nums) {
			temp = make([]int, 0, 3)
			if j >= len(nums)-2 {
				i++
				j = i + 1
				k = j + 1
			} else {
				j++
				k = j + 1
			}
			continue
		}

		if j == len(nums)-1 {
			temp = make([]int, 0, 3)
			i++
			j = i + 1
			k = j + 1
			continue
		}

		if len(temp) == 0 {
			temp = append(temp, nums[i])
		}

		if len(temp) == 1 {
			one := nums[j]
			temp = append(temp, one)
		}

		if len(temp) == 2 {
			two := nums[k]
			if temp[0]+temp[1]+two == 0 && (two != nums[k-1] || two == 0) {
				triplet := []int{temp[0], temp[1], two}
				res = append(res, triplet)
				temp = make([]int, 0, 3)
				//k++
			}
			k++
		}

	}
	return res
}

func filterNums(nums []int) []int {
	sort.Ints(nums)

	filtered := make([]int, 0, len(nums))

	filtered = append(filtered, nums[0])
	filtered = append(filtered, nums[1])
	zerosAmnt := 0
	for _, v := range filtered {
		if v == 0 {
			zerosAmnt++
		}
	}

	j := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] == 0 {
			if zerosAmnt < 3 {
				filtered = append(filtered, 0)
				zerosAmnt++
				j++
			}
			continue
		}
		cur := nums[i]
		if cur == filtered[j-1] && cur == filtered[j-2] {
			continue
		}
		j++
		filtered = append(filtered, cur)
	}

	return filtered
}
