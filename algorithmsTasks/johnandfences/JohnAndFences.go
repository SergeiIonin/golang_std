package main

import (
	"fmt"
)

// todo use sort.Ints(arr) instead!
type fence []int

func (f fence) Len() int {
	return len(f)
}
func (f fence) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f fence) Less(i, j int) bool {
	return f[i] < f[j]
}

func main() {
	arrayInt := []int{1, 54, 135, 8, 0, 89}
	maxElem := max(arrayInt)
	fmt.Println("max elem = ", maxElem)
	testFence0 := []int{1370, 4873, 2981, 478, 4760, 5191, 6872, 6665, 3327, 3106, 9828, 9991, 208, 1667, 8408, 6876, 4872, 320, 1675, 747,
		7706, 4165, 1579, 2988, 1126, 2093, 1313, 5300, 2111, 6948, 6838, 9833, 1821, 6171, 310, 2932, 7713, 3533, 9596, 1039, 6639, 5775,
		1030, 3198, 7441, 5789, 6425, 8665, 6108, 8099, 9411, 3814, 8616, 989, 6801, 9741, 9433, 4465, 5040, 1544, 1412, 8230, 7728, 3232,
		4400, 4389, 2515, 8464, 7922, 8463, 9503, 912, 589, 532, 461, 4382, 6320, 6885, 3046, 2427, 1335, 8808, 2592, 6302, 6149, 5744, 6043,
		5581, 208, 7434, 3476, 1620, 2015, 7555, 1203, 2766, 1944, 3718, 1230, 6217}
	res0 := getMaxSquare(testFence0, int64(0))
	fmt.Println("res = ", res0)
	fmt.Println("--------")

	testFence1 := []int{2, 5, 7, 4, 1, 8} // 12
	res1 := getMaxSquare(testFence1, int64(0))
	fmt.Println("res1 = ", res1)
	fmt.Println("--------")

	testFence2 := []int{0, 5, 7, 4, 6, 8} // 20
	res2 := getMaxSquare(testFence2, int64(0))
	fmt.Println("res2 = ", res2)
	fmt.Println("--------")
}

func getMaxSquare(f fence, square int64) int64 {
	if len(f) == 0 {
		return square
	} else if len(f) == 1 {
		fenceHead := f[0]
		if int64(fenceHead) >= square {
			return int64(fenceHead)
		} else {
			return square
		}
	} else if int64(len(f)*max(f)) <= square {
		return square
	} else {
		min := min(f)
		indexMin := indexOf(f, min)
		fmt.Println("indexMin = ", indexMin)

		squareNew := int64(min * len(f))

		fenceLeft := f[:indexMin] // [2 4 6 0 3 6 9]
		fmt.Println("fenceLeft = ", fenceLeft)
		fenceRight := f[indexMin+1:]
		fmt.Println("fenceRight = ", fenceRight)
		if square < squareNew {
			square = squareNew
		}
		fmt.Println("fence square max so far = ", square)
		left := getMaxSquare(fenceLeft, square)
		right := getMaxSquare(fenceRight, square)
		if left < right {
			return right
		} else {
			return left
		}
	}
}

func max(fence []int) int {
	len := len(fence)
	maxElem := fence[0]
	elem := maxElem
	for i := 1; i < len; i++ {
		elem = fence[i]
		if elem > maxElem {
			maxElem = elem
		}
	}
	return maxElem
}

func min(fence []int) int {
	len := len(fence)
	minElem := fence[0]
	elem := minElem
	for i := 1; i < len; i++ {
		elem = fence[i]
		if elem < minElem {
			minElem = elem
		}
	}
	return minElem
}

func indexOf(f fence, elem int) int {
	index := -1
	for i, current := range f {
		if current == elem {
			index = i
		}
	}
	if index == -1 {
		panic("elem wasn't found")
	}
	return index
}
