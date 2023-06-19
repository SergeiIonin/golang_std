package main

import "fmt"

// https://leetcode.com/problems/koko-eating-bananas/

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
// If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h hours.

func main() {
	in0 := []int{3, 6, 7, 11}
	h0 := 8
	r0 := minEatingSpeed(in0, h0)
	fmt.Println("r0 = ", r0) // 4

	in1 := []int{30, 11, 23, 4, 20}
	h1 := 5
	r1 := minEatingSpeed(in1, h1)
	fmt.Println("r1 = ", r1) // 30

	in2 := []int{30, 11, 23, 4, 20}
	h2 := 6
	r2 := minEatingSpeed(in2, h2)
	fmt.Println("r2 = ", r2) // 23

	in3 := []int{312884470}
	h3 := 312884469
	r3 := minEatingSpeed(in3, h3)
	fmt.Println("r3 = ", r3) // 2

	in4 := []int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}
	h4 := 823855818
	r4 := minEatingSpeed(in4, h4)
	fmt.Println("r4 = ", r4) // 14

	in5 := []int{1, 1, 1, 999999999}
	h5 := 10
	r5 := minEatingSpeed(in5, h5)
	fmt.Println("r5 = ", r5) // 142857143
}

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 0 {
		return 0
	}
	if len(piles) == 1 {
		if h > piles[0] {
			return 1
		} else {
			t := piles[0] / h
			if piles[0]%h != 0 {
				t += 1
			}
		}
	}
	minRate := 1
	maxRate := max(piles)
	res := (minRate + maxRate) / 2
	delta := h

	for minRate <= maxRate {
		rate := (minRate + maxRate) / 2
		t := getTime(piles, rate)
		deltaNext := h - t
		if deltaNext >= 0 && deltaNext <= delta {
			delta = deltaNext
			res = rate
		}
		if t > h {
			minRate = rate + 1
		} else if t <= h {
			maxRate = rate - 1
		}
	}

	return res

}

func getTime(piles []int, rate int) int {
	t := 0
	for _, p := range piles {
		if p <= rate {
			t += 1
		} else {
			t += p / rate
			if p%rate != 0 {
				t += 1
			}
		}
	}
	return t
}

func max(piles []int) int {
	m := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > m {
			m = piles[i]
		}
	}
	return m
}

// func min(piles []int) int {
// 	m := piles[0]
// 	for i := 1; i < len(piles); i++ {
// 		if piles[i] < m {
// 			m = piles[i]
// 		}
// 	}
// 	return m
// }
