package main

import (
	"fmt"
)

func main() {
	arrayInt := []int{1, 54, 135, 8, 0, 89}
	maxElem := max(arrayInt)
	fmt.Println("max elem = ", maxElem)
	testHeights0 := []int{1370, 4873, 2981, 478, 4760, 5191, 6872, 6665, 3327, 3106, 9828, 9991, 208, 1667, 8408, 6876, 4872, 320, 1675, 747,
		7706, 4165, 1579, 2988, 1126, 2093, 1313, 5300, 2111, 6948, 6838, 9833, 1821, 6171, 310, 2932, 7713, 3533, 9596, 1039, 6639, 5775,
		1030, 3198, 7441, 5789, 6425, 8665, 6108, 8099, 9411, 3814, 8616, 989, 6801, 9741, 9433, 4465, 5040, 1544, 1412, 8230, 7728, 3232,
		4400, 4389, 2515, 8464, 7922, 8463, 9503, 912, 589, 532, 461, 4382, 6320, 6885, 3046, 2427, 1335, 8808, 2592, 6302, 6149, 5744, 6043,
		5581, 208, 7434, 3476, 1620, 2015, 7555, 1203, 2766, 1944, 3718, 1230, 6217}
	res0 := maxArea(testHeights0)
	fmt.Println("res = ", res0)
	fmt.Println("--------")

	testHeights1 := []int{2, 5, 7, 4, 1, 8} // 12
	res1 := maxArea(testHeights1)
	fmt.Println("res1 = ", res1)
	fmt.Println("--------")

	testHeights2 := []int{0, 5, 7, 4, 6, 8} // 20
	res2 := maxArea(testHeights2)
	fmt.Println("res2 = ", res2)
	fmt.Println("--------")
}

// https://leetcode.com/problems/container-with-most-water/description/

func maxArea(height []int) int {
	return int(getMaxSquare(height, 0))
}

type height []int

func (h height) Len() int {
	return len(h)
}
func (h height) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h height) Less(i, j int) bool {
	return h[i] < h[j]
}

func getMaxSquare(h height, square int) int {
	if len(h) == 0 {
		return square
	} else if len(h) == 1 {
		head := h[0]
		if head >= square {
			return head
		} else {
			return square
		}
	} else if len(h)*max(h) <= square {
		return square
	} else {
		min := min(h)
		indexMin := indexOf(h, min)
		fmt.Println("indexMin = ", indexMin)

		squareNew := min * len(h)

		hLeft := h[:indexMin]
		fmt.Println("hLeft = ", hLeft)
		hRight := h[indexMin+1:]
		fmt.Println("hRight = ", hRight)
		if square < squareNew {
			square = squareNew
		}
		fmt.Println("h square max so far = ", square)
		left := getMaxSquare(hLeft, square)
		right := getMaxSquare(hRight, square)
		if left < right {
			return right
		} else {
			return left
		}
	}
}

func max(h []int) int {
	len := len(h)
	maxElem := h[0]
	elem := maxElem
	for i := 1; i < len; i++ {
		elem = h[i]
		if elem > maxElem {
			maxElem = elem
		}
	}
	return maxElem
}

func min(h []int) int {
	len := len(h)
	minElem := h[0]
	elem := minElem
	for i := 1; i < len; i++ {
		elem = h[i]
		if elem < minElem {
			minElem = elem
		}
	}
	return minElem
}

func indexOf(h height, elem int) int {
	index := -1
	for i, current := range h {
		if current == elem {
			index = i
		}
	}
	if index == -1 {
		panic("elem wasn't found")
	}
	return index
}
