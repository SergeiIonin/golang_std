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

	fmt.Println("=======")

}

func trap_(height []int) int {

	if len(height) < 3 {
		return 0
	}

	v := 0

	i := 0
	j := 1

	for i < len(height)-1 {
		if j == len(height) {
			v += getVolume(height[i:])
			// if height[i] > height[j-1] {
			// 	break
			// }
			i++
			j = i + 1
			continue
		}
		if height[j] >= height[i] {
			if j-i > 1 {
				v += getVolume(height[i : j+1])
			}
			i = j
			j = i + 1
		} else {
			j++
		}
	}
	return v
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
	for i < j {
		heightsInv[i] = heights[j]
		i++
		j--
	}
	return heightsInv
}

func volLast(heights []int) int {
	heightsInv := make([]int, len(heights))
	i := 0
	j := len(heights) - 1
	for i < j {
		heightsInv[i] = heights[j]
		i++
		j--
	}
	return trap(heightsInv)
}

func trap__(height []int) int {

	if len(height) < 3 {
		return 0
	}

	v := 0
	interval := []int{}
	interval = append(interval, height[0])
	fall := true
	i := 0
	if height[0] == 0 {
		height = height[1:]
	}
	if height[len(height)-1] == 0 {
		height = height[:len(height)-1]
	}

	for i < len(height) {
		j := i + 1
		if j < len(height) {
			if fall {
				if height[j] > height[i] {
					fall = false
				}
			} else { // rise
				if height[j] >= interval[0] {
					interval = append(interval, height[j])
					i = j
					v += getVolume(interval)
					fall = true
					interval = append([]int{}, height[i])
					continue
				} else if height[j] <= height[i] {
					v += getVolume(interval)
					fall = true
					interval = append([]int{}, height[i])
					continue
				}
			}
			interval = append(interval, height[j])
		}
		i++
	}
	return v
}

func getVol(m int, heights []int) int {
	v := m * len(heights)
	for _, h := range heights {
		v = v - h
	}
	return v
}

func minAbs(l, r int) int {
	aL, aR := abs(l), abs(r)
	if aL < aR {
		return aL
	} else {
		return aR
	}
}

func abs(a int) int {
	if a < 0 {
		return a
	} else {
		return -a
	}
}

func getVolume(heights []int) int {
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
	if v < 0 {
		v = 0
	}
	fmt.Println("v = ", v)
	return v
}

func min(l, r int) int {
	if l < r {
		return l
	} else {
		return r
	}
}
