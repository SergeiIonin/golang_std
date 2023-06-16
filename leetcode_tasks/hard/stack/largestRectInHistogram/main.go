package main

import "fmt"

// https://leetcode.com/problems/largest-rectangle-in-histogram/

// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1,
// return the area of the largest rectangle in the histogram.

func main() {
	in0 := []int{3, 2, 1, 2, 3, 4}
	r0 := largestRectangleArea(in0)
	fmt.Println(r0) // 6

	in1 := []int{2, 1, 5, 6, 2, 3}
	r1 := largestRectangleArea(in1)
	fmt.Println(r1) // 10

	in2 := []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	r2 := largestRectangleArea(in2)
	fmt.Println(r2) // 12
}

type StackValue struct {
	index  int
	height int
}

// neetcode solution https://github.com/neetcode-gh/leetcode/blob/main/go/0084-largest-rectangle-in-histogram.go
func largestRectangleArea(heights []int) int {
	stack := []StackValue{} // pair: {index, height}
	maxArea := 0
	var start int

	for i, h := range heights {
		start = i
		for len(stack) != 0 && stack[len(stack)-1].height > h {
			index, height := stack[len(stack)-1].index, stack[len(stack)-1].height
			stack = stack[0 : len(stack)-1] //pop top from stack
			maxArea = max(maxArea, height*(i-index))
			start = index
		}
		stack = append(stack, StackValue{start, h})
	}

	for _, h := range stack {
		maxArea = max(maxArea, h.height*(len(heights)-h.index))
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
