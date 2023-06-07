package main

import "fmt"

func main() {
	res := generate(6)
	fmt.Println(res)
}

func generate(numRows int) [][]int {
	accum := make([][]int, numRows)
	accum[0] = []int{1}
	return rec(1, accum)
}

func rec(lvl int, accum [][]int) [][]int {
	if lvl == len(accum) {
		return accum
	} else {
		prev := accum[lvl-1]
		lenNext := len(prev) + 1
		next := make([]int, lenNext)
		for i := 0; i < lenNext; i++ {
			if i == 0 || i == lenNext-1 {
				next[i] = 1
			} else {
				next[i] = prev[i-1] + prev[i]
			}
		}
		accum[lvl] = next
		return rec(lvl+1, accum)
	}
}
