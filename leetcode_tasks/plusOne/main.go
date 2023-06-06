package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{9, 8, 9}))
	fmt.Println(plusOne([]int{3, 3, 2}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}

// this solution is less performant:
// func plusOne(digits []int) []int {
// 	for i := len(digits) - 1; i >= 0; i-- {
// 		if digits[i] == 9 {
// 			digits[i] = 0
// 			if i == 0 {
// 				digits = append([]int{1}, digits...)
// 				return digits
// 			}
// 		} else {
// 			digits[i]++
// 			return digits
// 		}
// 	}
// 	return digits
// }
