package main

import "fmt"

func main() {
	in0 := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	t0_0 := 3
	t0_1 := 63
	t0_2 := -1

	r0_0 := searchMatrix(in0, t0_0)
	r0_1 := searchMatrix(in0, t0_1)
	r0_2 := searchMatrix(in0, t0_2)

	fmt.Println(r0_0)
	fmt.Println(r0_1)
	fmt.Println(r0_2)
}

func searchMatrix(matrix [][]int, target int) bool {
	rowsMax := len(matrix) - 1
	colsMax := len(matrix[0]) - 1

	rD, rU := 0, rowsMax

	for rD <= rU {
		rM := (rD + rU) / 2

		if target < matrix[rM][0] {
			rU = rM - 1
		} else if target > matrix[rM][colsMax] {
			rD = rM + 1
		} else {
			return searchRow(matrix[rM], target)
		}
	}

	return false

}

func searchRow(row []int, target int) bool {
	left := 0
	right := len(row) - 1

	for left <= right {
		m := (left + right) / 2

		if row[m] == target {
			return true
		} else if target < row[m] {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return false
}
