package main

import "fmt"

func main() {

	in0 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	r0 := dailyTemperatures(in0)
	fmt.Println(r0) // [1,1,4,2,1,1,0,0]

	in1 := []int{30, 40, 50, 60}
	r1 := dailyTemperatures(in1)
	fmt.Println(r1) // [1,1,1,0]

	in2 := []int{30, 60, 90}
	r2 := dailyTemperatures(in2)
	fmt.Println(r2) // [1,1,0]

}

// based on the most performant solution on leetcode (but different in that I manually incr i in the cycle)
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	i := 0
	for i < len(temperatures) {
		if len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			res[prev] = i - prev
			stack = stack[:len(stack)-1]
			continue
		} else {
			stack = append(stack, i)
			i++
		}
	}

	return res
}
