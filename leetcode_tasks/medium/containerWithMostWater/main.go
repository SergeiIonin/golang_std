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

	fmt.Println("a0 = ", a0) // 49
	fmt.Println("a1 = ", a1) // 1
	fmt.Println("a2 = ", a2) // 16
	fmt.Println("a3 = ", a3) // 15
	fmt.Println("a4 = ", a4) // 4913370
	fmt.Println("a5 = ", a5) // 721777500

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

func maxArea_(height []int) int {

	s := 0
	sTemp := 0

	i := 0
	j := len(height) - 1

	for i < j {
		l := height[i]
		r := height[j]
		len := (j - i)
		if l > r {
			sTemp = len * r
			j--
		} else {
			sTemp = len * l
			i++
		}
		if sTemp > s {
			s = sTemp
		}
	}

	return s
}

func maxArea(height []int) int {
	s := 0

	i := 0
	j := len(height) - 1

	for i <= j {
		sCur := (j - i) * min(height[i], height[j])
		if sCur > s {
			s = sCur
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return s
}

func min(l, r int) int {
	if l > r {
		return r
	} else {
		return l
	}
}
