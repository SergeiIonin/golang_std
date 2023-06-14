package main

import (
	"fmt"
)

// https://leetcode.com/problems/trapping-rain-water/

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

func main() {

	in0 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	r0 := trap(in0)
	fmt.Println(r0) // 6

	fmt.Println("=======")

	in1 := []int{4, 2, 0, 3, 2, 5}
	r1 := trap(in1)
	fmt.Println(r1) // 9

	fmt.Println("=======")

	in2 := []int{2, 0, 2}
	r2 := trap(in2)
	fmt.Println(r2) // 2

	fmt.Println("=======")

	in3 := []int{4, 2, 3}
	r3 := trap(in3)
	fmt.Println(r3) // 1

	fmt.Println("=======")

	in4 := []int{9, 8, 2, 6}
	r4 := trap(in4)
	fmt.Println(r4) // 4

	in5 := []int{1, 2, 1, 2, 3}
	r5 := trap(in5)
	fmt.Println(r5) // 1

	fmt.Println("=======")

}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	v := 0
	i := 0

	for i < len(height) {
		in := append([]int{}, height[i])
		for j := i + 1; j < len(height); j++ {
			in = append(in, height[j])
			if height[j] >= height[i] {
				v += vol(in)
				i = j
				break
			}
			if j == len(height)-1 {
				v += trap(reverse(in))
				i = j
				break
			}
		}
		if i == len(height)-1 {
			break
		}
	}

	return v

}

func vol(heights []int) int {
	v := 0
	if len(heights) < 3 {
		return 0
	} else {
		l := heights[0]
		r := heights[len(heights)-1]
		m := min(l, r)
		heightsShort := heights[1 : len(heights)-1]
		v = m * len(heightsShort)
		for _, h := range heightsShort {
			v = v - h
		}
	}
	// if v < 0 {
	// 	v = volLast(heights)
	// }
	// fmt.Println("v = ", v)
	return v
}

func reverse(heights []int) []int {
	heightsInv := make([]int, len(heights))
	i := 0
	j := len(heights) - 1
	for i < len(heights) && j >= 0 {
		heightsInv[i] = heights[j]
		i++
		j--
	}
	return heightsInv
}

func min(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
