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

/*
https://leetcode.com/problems/container-with-most-water
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
Find two lines that together with the x-axis form a container, such that the container contains the most water.
Return the maximum amount of water a container can store.
Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1


Constraints:
    n == height.length
    2 <= n <= 10^5
    0 <= height[i] <= 10^4
*/
func maxArea(height []int) int {
    left := 0
	right := len(height)-1

	min := func(a, b int) int {
		if a <= b {
			return a
		} else {
			return b
		}
	}

	MAX := 10_000

	cond := func(volMax, base int) bool {
		return base * MAX < volMax
	}

	vol := 0

	for ; left != right ; {
		base := right - left
		if cond(vol, base) {
			return vol
		}
		hL := height[left]
		hR := height[right]
		vol_cur := min(hL, hR) * (right - left)
		if vol_cur > vol {
			vol = vol_cur
		}
		if hL <= hR {
			left++
		} else {
			right--
		}
	}
	return vol
}
