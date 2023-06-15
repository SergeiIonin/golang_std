package main

import "fmt"

// https://leetcode.com/problems/largest-rectangle-in-histogram/

// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
// return the area of the largest rectangle in the histogram.

func main() {
	// in0 := []int{3, 2, 1, 2, 3, 4}
	// r0 := largestRectangleArea(in0)
	// fmt.Println(r0) // 6

	// in1 := []int{2, 1, 5, 6, 2, 3}
	// r1 := largestRectangleArea(in1)
	// fmt.Println(r1) // 10

	in2 := []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	r2 := largestRectangleArea(in2)
	fmt.Println(r2) // 12
	// 5,5,1,7,1,1,5,2,7,6
}

func largestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights))
	occurences := make([]int, len(heights))

	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			occurences[i]++
		} else if heights[i] < heights[stack[len(stack)-1]] {

			n := 0
			for j := len(stack) - 1; j >= 0; j-- {
				index := stack[j]
				occurences[index] += (len(stack) - j - 1)
				if heights[index] > heights[i] {
					n++
				}
			}

			occurences[i] = 1
			for j := i - 1; j >= 0; j-- {
				if heights[j] >= heights[i] {
					occurences[i]++
				} else {
					break
				}
			}

			//stack = stack[:len(stack)-n]
			stack = append([]int{}, i)
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		index := stack[i]
		occurences[index] += (len(stack) - i - 1)
	}

	max := 0

	for i := 0; i < len(heights); i++ {
		s := heights[i] * occurences[i]
		if s > max {
			max = s
		}
	}

	return max
}
