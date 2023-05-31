package main

import (
	"fmt"
)

func main() {
	input0 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // 49
	input1 := []int{1, 1}                      // 1
	input2 := []int{1, 5, 16, 16, 5, 3}        // 16
	input3 := []int{1, 5, 16, 14, 5, 3}        // 15

	a0 := maxArea(input0)
	a1 := maxArea(input1)
	a2 := maxArea(input2)
	a3 := maxArea(input3)

	fmt.Println("a0 = ", a0)
	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
	fmt.Println("a3 = ", a3)

}

// https://leetcode.com/problems/container-with-most-water/description/

// h = [1,8,6,2,5,4,8,3,7]
// 		sorted0 = 8
//		sorted1 = 8
//		diff = 5
//		s = min(sorted0, sorted1) * diff // 40
//
//

func maxArea(height []int) int {

	size := len(height)
	visitedIndexes := make([]int, 0, size)
	s := 0

	for i := 0; i < size; i++ {
		max, index := maxWithIndex(height, &visitedIndexes)
		fmt.Println("max = ", max)
		fmt.Println("visitedIndexes = ", visitedIndexes)
		if max*size < s {
			continue
		}
		s = getS(height, max, index, s, size)
	}

	return s
}

func indexVisited(ind int, indexes []int) bool {
	for i := 0; i < len(indexes); i++ {
		if indexes[i] == ind {
			return true
		}
	}
	return false
}

// current max among the elems which haven't been visited yet
// it works but it's not enough fast
func maxWithIndex(h []int, visitedIndexes *[]int) (max int, index int) {
	index = 0
	max = h[index]
	size := len(h)
	if size == 1 {
		return
	}
	for i := 0; i < size; i++ {
		if indexVisited(i, *visitedIndexes) {
			continue
		}
		if max < h[i] {
			max = h[i]
			index = i
		}
	}
	*visitedIndexes = append(*visitedIndexes, index)
	return
}

func minOfTwo(left int, right int) int {
	if left < right {
		return left
	} else {
		return right
	}
}

func getS(h []int, max int, index int, sMax int, size int) (s int) {
	sTemp := 0
	s = sMax
	fmt.Println("sBase = ", s)
	for i, v := range h {
		if index == i {
			continue
		}
		sTemp = (index - i) * minOfTwo(max, v)
		if sTemp < 0 {
			sTemp = sTemp * (-1)
		}
		if sTemp > s {
			s = sTemp
		}
	}
	fmt.Println("s = ", s)
	return s
}
