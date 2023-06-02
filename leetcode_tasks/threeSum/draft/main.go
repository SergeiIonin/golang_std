package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	size := len(nums)

	if size < 3 {
		return [][]int{}
	}

	res := make([][]int, 0, size)

	sort.Ints(nums)

	max := nums[size-1]

	ind_pos := 0

	zeros := make([]int, 0, 3)

	for i := 0; i < size; i++ {
		if nums[i] == 0 && len(zeros) < 3 {
			zeros = append(zeros, 0)
		}
		if nums[i] > 0 {
			ind_pos = i
			break
		}
	}

	pos := make([]int, 0)
	min_pos := 0
	if ind_pos > 0 {
		pos = nums[ind_pos:]
		min_pos = pos[0]
	}

	sizePos := len(pos)

	if len(zeros) == 3 {
		res = append(res, []int{0, 0, 0})
	}

	i := 0
	j := 1

	for i < size-2 {

		zero := nums[i]
		if zero > 0 {
			break
		}
		one := nums[j]

		if j == size-1 {
			if i >= size-3 {
				break
			}
			i++
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
		if two > max {
			j++
			continue
		}
		if two < min_pos && two > 0 {
			j++
			continue
		}
		for k := 0; k < sizePos; k++ {
			elem := pos[k]
			if elem <= one {
				continue
			}
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
