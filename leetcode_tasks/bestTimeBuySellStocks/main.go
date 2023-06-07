package main

import "fmt"

func main() {
	p0 := []int{7, 1, 5, 3, 6, 4}
	p1 := []int{7, 6, 4, 3, 1}

	r0 := maxProfit(p0)
	r1 := maxProfit(p1)

	fmt.Println(r0)
	fmt.Println(r1)

}

// 56.62% RT (129 ms), 8.1 MB (82.81%)
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	min := prices[0]
	next := 0
	diff := 0
	diffMax := 0

	for i := 1; i < len(prices); i++ {
		next = prices[i]
		diff = next - min
		if diff > diffMax {
			diffMax = diff
		}
		if next < min {
			min = next
		}
	}
	return diffMax
}
