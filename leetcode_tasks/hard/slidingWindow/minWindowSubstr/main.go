package main

import "fmt"

// https://leetcode.com/problems/minimum-window-substring/

// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring
//  of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

// The testcases will be generated such that the answer is unique.

func main() {
	// s := "ADOBECODEBANC"
	// t := "ABC"
	s := "a"
	t := "a"
	r := minWindow(s, t)
	fmt.Println(r)
}

func minWindow(s string, t string) string {
	runes_s := []rune(s)
	runes_t := []rune(t)

	hashmap_t := make(map[rune]int, len(runes_t))

	hashmap_res := make(map[rune]int, len(runes_t))

	for _, r := range runes_t {
		hashmap_t[r]++
		hashmap_res[r] = 0
	}

	i := 0
	j := len(runes_t) - 1
	start := 0
	end := len(runes_s) + 1
	size := len(runes_s) + 1
	substr := runes_s[i : j+1]

	for _, r := range substr {
		if _, ok := hashmap_t[r]; ok {
			hashmap_res[r]++
		}
	}

	for j < len(runes_s) && i <= j {
		if isFull(hashmap_res, hashmap_t) {
			if (j + 1 - i) < size {
				start = i
				end = j + 1
				size = end - start
				if size == len(runes_t) {
					break
				}
			}
			if _, ok := hashmap_res[runes_s[i]]; ok {
				hashmap_res[runes_s[i]]--
			}
			k := i + 1
			for ; k < j; k++ {
				if _, ok := hashmap_t[runes_s[k]]; ok {
					break
				}
			}
			i = k
		} else {
			k := j + 1
			for ; k < len(runes_s); k++ {
				if _, ok := hashmap_t[runes_s[k]]; ok {
					hashmap_res[runes_s[k]]++
					break
				}
			}
			j = k
		}
	}

	if end <= len(runes_s) {
		return string(runes_s[start:end])
	} else {
		return ""
	}

}

func isFull(hashmap map[rune]int, base map[rune]int) bool {
	for r, occ := range hashmap {
		if occ < base[r] {
			return false
		}
	}
	return true
}

// func minWindow(s string, t string) string {
// 	runes_s := []rune(s)
// 	runes_t := []rune(t)

// 	hashmap_s := make(map[rune][]int, len(runes_s))
// 	hashmap_t := make(map[rune]int, len(runes_t))

// 	for _, r := range runes_t {
// 		hashmap_t[r]++
// 	}

// 	for i, r := range runes_s {
// 		if _, ok := hashmap_t[r]; ok {
// 			hashmap_s[r] = append(hashmap_s[r], i)
// 		}
// 	}

// 	i := 0
// 	j := len(runes_t) - 1
// 	start := 0
// 	end := len(runes_s) + 1
// 	size := len(runes_s) + 1

// 	for j < len(runes_s) {
// 		if containsAllElems(hashmap_t, hashmap_s, i, j+1) {
// 			if (j + 1 - i) < size {
// 				start = i
// 				end = j + 1
// 				size = j + 1 - i
// 				if size == len(runes_t) {
// 					break
// 				}
// 			}

// 			k := i + 1
// 			for ; k < len(runes_s); k++ {
// 				if _, ok := hashmap_t[runes_s[k]]; ok {
// 					break
// 				}
// 			}
// 			i = k

// 		} else {
// 			k := j + 1
// 			for ; k < len(runes_s); k++ {
// 				if _, ok := hashmap_t[runes_s[k]]; ok {
// 					break
// 				}
// 			}
// 			j = k
// 		}
// 	}
// 	if end <= len(runes_s) {
// 		return string(runes_s[start:end])
// 	} else {
// 		return ""
// 	}
// }

// func containsAllElems(hashmap_t map[rune]int, hashmap_s map[rune][]int, start int, end int) bool {
// 	for r, occ := range hashmap_t {
// 		if slice, ok := hashmap_s[r]; ok && len(slice) >= occ {
// 			count := occ
// 			for _, ind := range slice {
// 				if ind >= start && ind < end {
// 					count--
// 					if count == 0 {
// 						break
// 					}
// 				}
// 			}
// 			if count > 0 {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }
