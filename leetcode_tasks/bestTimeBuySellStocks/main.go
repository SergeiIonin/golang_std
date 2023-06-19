package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

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

// this solution from leetcode is really good:
// func maxProfit(prices []int) int {
//     res, min := 0, prices[0]
//     for i:=1;i<len(prices);i++ {
//         res = maxV(res, prices[i]-min)
//         min = minV(min, prices[i])
//     }
//     return res
// }

// func minV(i, j int) int {
//     if i<j{
//         return i
//     }else {
//         return j
//     }
// }

// func maxV(i, j int) int {
//     if i>j{
//         return i
//     }else {
//         return j
//     }
// }
