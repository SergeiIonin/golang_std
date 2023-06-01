package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	input0 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // 49
	input1 := []int{1, 1}                      // 1
	input2 := []int{1, 5, 16, 16, 5, 3}        // 16
	input3 := []int{1, 5, 16, 14, 5, 3}        // 15
	f0, err := os.Open("resources/input")
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
	f2, err := os.Open("resources/input2")
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
	input4 := readInput(f0) // 4913370
	input5 := readInput(f2) // 721777500

	start := time.Now()

	a0 := maxArea(input0)
	a1 := maxArea(input1)
	a2 := maxArea(input2)
	a3 := maxArea(input3)
	a4 := maxArea(input4)
	a5 := maxArea(input5)

	timeElapsed := time.Since(start)
	fmt.Printf("Time elapsed = %d \n", timeElapsed.Microseconds())

	fmt.Println("a0 = ", a0)
	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
	fmt.Println("a3 = ", a3)
	fmt.Println("a4 = ", a4)
	fmt.Println("a5 = ", a5)

}

func readInput(f *os.File) []int {
	r := bufio.NewReader(f)
	ints := make([]int, 0, 10000)

	for {
		b, err := r.ReadBytes(byte(','))
		if err != nil {
			fmt.Println("end of file:", err.Error())
			break
		}
		bNew := bytes.TrimSuffix(b, []byte(","))

		s := string(bNew)
		int, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err.Error())
		}
		ints = append(ints, int)
	}
	return ints
}

// https://leetcode.com/problems/container-with-most-water/description/

// h = [1,8,6,2,5,4,8,3,7]
// 		sorted0 = 8
//		sorted1 = 8
//		diff = 5
//		s = min(sorted0, sorted1) * diff // 40
//
//

// func maxArea(height []int) int {

// 	size := len(height)
// 	s := 0

// 	var lenLeft, lenRight int

// 	for ind, elem := range height {
// 		if elem == 0 {
// 			continue
// 		}

// 		lenLeft = ind
// 		lenRight = size - lenLeft - 1

// 		if lenLeft*elem < s && lenRight*elem < s {
// 			continue
// 		}

// 		sCur := 0

// 		for j, cur := range height {
// 			if cur < elem || ind == j {
// 				continue
// 			}
// 			sCur = abs(ind-j) * elem
// 			if sCur > s {
// 				s = sCur
// 			}
// 		}
// 	}

// 	return s

// }

// func abs(v int) int {
// 	if v < 0 {
// 		return -v
// 	}
// 	return v
// }

func maxArea(height []int) int {

	size := len(height)
	s := 0

	for i, elem := range height {
		if i == len(height)-1 {
			break
		}

		if elem == 0 {
			continue
		}

		if i*elem < s && (size-i-1)*elem < s {
			continue
		}

		sCur := 0
		for j := (i + 1); j < len(height); j++ {
			cur := height[j]
			if cur == 0 {
				continue
			}
			h := 0
			if cur < elem {
				h = cur
			} else {
				h = elem
			}
			sCur = (j - i) * h
			if sCur > s {
				s = sCur
			}
		}
	}

	return s
}
