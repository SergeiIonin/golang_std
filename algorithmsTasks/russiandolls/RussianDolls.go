package main

import (
	"fmt"
	"sort"
)

func main() {
	testDolls0 := []int{2, 2, 3, 3}                      // 2
	testDolls1 := []int{1, 2, 2, 3, 4, 5, 6}             // 2
	testDolls2 := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 5} // 3
	testDolls3 := []int{1, 2, 3}                         // 1
	testDolls4 := []int{}
	res0 := getRemainingDolls(testDolls0)
	fmt.Println("res0 = ", res0)
	res1 := getRemainingDolls(testDolls1)
	fmt.Println("res1 = ", res1)
	res2 := getRemainingDolls(testDolls2)
	fmt.Println("res2 = ", res2)
	res3 := getRemainingDolls(testDolls3)
	fmt.Println("res3 = ", res3)
	res4 := getRemainingDolls(testDolls4)
	fmt.Println("res4 = ", res4)
}

type dolls []int

func (d dolls) Len() int {
	return len(d)
}
func (d dolls) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d dolls) Less(i, j int) bool {
	return d[i] < d[j]
}

func getRemainingDolls(d dolls) int {
	sort.Sort(d)
	return getTotalRemainedDolls(d, 0)
}

func getTotalRemainedDolls(d dolls, rem int) int {
	if len(d) == 0 {
		return rem
	} else if len(d) == 1 {
		return rem + 1
	} else {
		remainedDolls := remainder(d, []int{})
		fmt.Println("remainedDolls = ", remainedDolls)
		return getTotalRemainedDolls(remainedDolls, rem+1)
	}
}

func remainder(d dolls, rem dolls) dolls {
	if len(d) <= 1 {
		return rem
	} else if d[0] < d[1] {
		return remainder(d[1:], rem)
	} else {
		if len(d) >= 3 {
			rem = append(rem, d[1])
			d = append([]int{d[0]}, d[2:]...)
		} else if len(d) == 2 {
			rem = append(rem, d[1])
			d = []int{d[0]}
		}
		return remainder(d, rem)
	}
}
