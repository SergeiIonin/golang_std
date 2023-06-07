package main

import "fmt"

func main() {
	res := generate(6)
	fmt.Println(res)
}

// this solution beats 100% by runtime and 41.94% by memory
func generate(numRows int) [][]int {
	accum := make([][]int, numRows)
	accum[0] = []int{1}
	for i := 1; i < len(accum); i++ {
		prev := accum[i-1]
		lenNext := len(prev) + 1
		next := make([]int, lenNext)
		for j := 0; j < lenNext; j++ {
			if j == 0 || j == lenNext-1 {
				next[j] = 1
			} else {
				next[j] = prev[j-1] + prev[j]
			}
		}
		accum[i] = next
	}
	return accum
}
