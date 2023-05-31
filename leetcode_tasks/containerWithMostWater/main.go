package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input0 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7} // 49
	input1 := []int{1, 1}                      // 1
	input2 := []int{1, 5, 16, 16, 5, 3}        // 16
	input3 := []int{1, 5, 16, 14, 5, 3}        // 15
	input4 := readInput()                      // 4913370

	a0 := maxArea(input0)
	a1 := maxArea(input1)
	a2 := maxArea(input2)
	a3 := maxArea(input3)
	a4 := maxArea(input4)

	fmt.Println("a0 = ", a0)
	fmt.Println("a1 = ", a1)
	fmt.Println("a2 = ", a2)
	fmt.Println("a3 = ", a3)
	fmt.Println("a4 = ", a4)

}

func readInput() []int {
	f, err := os.Open("resources/input")
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
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

func maxArea(height []int) int {

	size := len(height)
	s := 0

	var heightWithIndexes HeightIndexed
	heightWithIndexes = make([]ElemWithIndex, size)

	for i, v := range height {
		heightWithIndexes[i] = ElemWithIndex{elem: v, index: i}
	}

	sort.Sort(heightWithIndexes) // asdending order
	lenLeft := 0
	lenRight := 0
	var elemWithIndex ElemWithIndex

	for i := 0; i < size; i++ {
		elemWithIndex = heightWithIndexes[i]
		el := elemWithIndex.elem
		ind := elemWithIndex.index

		lenLeft = elemWithIndex.index
		lenRight = size - lenLeft - 1

		if lenLeft*el < s && lenRight*el < s {
			break
		}

		sCur := 0

		for j, cur := range height {
			if cur < el || ind == j {
				continue
			}
			sCur = abs(ind-j) * el
			if sCur > s {
				s = sCur
			}
		}
	}

	return s

}

type ElemWithIndex struct {
	elem  int
	index int
}

type HeightIndexed []ElemWithIndex

func (hi HeightIndexed) Len() int {
	return len(hi)
}
func (hi HeightIndexed) Swap(i, j int) {
	hi[i], hi[j] = hi[j], hi[i]
}
func (hi HeightIndexed) Less(i, j int) bool {
	return hi[i].elem < hi[j].elem
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
