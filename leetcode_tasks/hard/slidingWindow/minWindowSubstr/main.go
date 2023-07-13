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
	s := "abc"
	t := "ac"
	r := minWindow(s, t)
	fmt.Println(r)
}

func minWindow(s string, t string) string {
	runes_s := []rune(s)
	runes_t := []rune(t)

	hashmap_t := make(map[rune]int, len(runes_t))
	hashmap_res := make(map[rune]int, len(runes_t))

	if len(runes_s) == 1 && len(runes_t) == 1 {
		if runes_s[0] == runes_t[0] {
			return s
		}
		return ""
	}

	for _, r := range runes_t {
		hashmap_t[r]++
		hashmap_res[r] = 0
	}

	i := -1
	valid := false
	j := 0

	// we fill the hashmap_res and start, end indexes
	for ; j < len(runes_s); j++ {
		if _, ok := hashmap_t[runes_s[j]]; ok {
			if i == -1 {
				i = j
			}
			hashmap_res[runes_s[j]]++
		}
		if (j-i+1) >= len(runes_t) && isFull(hashmap_res, hashmap_t) { // we can do better, e.g. countdown the elements in the copy of hashmap_t
			valid = true
			break
		}
	}

	start := i
	end := j + 1
	size := end - start

	// at this point hashmap_res is full, we should just keep it so
	for j < len(runes_s) && i <= j {

		removed := runes_s[i]
		hashmap_res[removed]--
		occ := hashmap_res[removed]
		if occ < hashmap_t[removed] { // if hashmap_res is no longer full, add elems
			k := j + 1
			found := false // indicator that the removed element is found
			for ; k < len(runes_s); k++ {
				if runes_s[k] == removed {
					hashmap_res[removed]++
					found = true
					break
				}
				if _, ok := hashmap_t[runes_s[k]]; ok {
					hashmap_res[runes_s[k]]++
				}
			}
			if found {
				j = k
			} else {
				break // if the removed rune isn't found, we should exit
			}
		}
		k := i + 1
		for ; k < j; k++ {
			if _, ok := hashmap_t[runes_s[k]]; ok { // move left index (i) to the next rune presented in the "t"
				break
			}
		}
		i = k

		if (j + 1 - i) < size {
			start = i
			end = j + 1
			size = end - start
			if size == len(runes_t) {
				break
			}
		}
	}

	if end <= len(runes_s) && valid {
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
