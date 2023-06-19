package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest
// substring without repeating characters.

func main() {
	in0 := "abcabcbb"  // 3
	in1 := "bbbbb"     // 1
	in2 := "pwwkew"    // 3
	in3 := "au"        // 2
	in4 := "bpfbhmipx" // 7
	in5 := "abba"      // 7

	r5 := lengthOfLongestSubstring(in5)
	r0 := lengthOfLongestSubstring(in0)
	r4 := lengthOfLongestSubstring(in4)
	r1 := lengthOfLongestSubstring(in1)
	r2 := lengthOfLongestSubstring(in2)
	r3 := lengthOfLongestSubstring(in3)

	fmt.Println(r0)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)
}

func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}
	hashmap := make(map[rune]int, len(runes))
	lenMax := 1
	start := 0

	for i, r := range runes {
		ind, ok := hashmap[r]
		if !ok {
			hashmap[r] = i
		} else {
			l := i - start
			if l > lenMax {
				lenMax = l
			}
			start = ind + 1
			hashmap[r] = i
		}
	}
	l := len(runes) - start
	if l > lenMax {
		lenMax = l
	}
	return lenMax
}
