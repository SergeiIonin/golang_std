package main

import "fmt"

// https://leetcode.com/problems/permutation-in-string/

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

// In other words, return true if one of s1's permutations is the substring of s2.

func main() {

	s1 := "ab"
	s2 := "eidboaooo"
	res := checkInclusion(s1, s2)

	fmt.Println(res)

}

func checkInclusion(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	if len(runes1) > len(runes2) {
		return false
	}

	s1Norm := normalize(runes1)
	hashmap := make(map[[26]byte]struct{}, 1)
	hashmap[s1Norm] = struct{}{}

	for i := 0; i < len(runes2)-len(runes1)+1; i++ {
		substr := runes2[i:(i + len(runes1))]
		substrNorm := normalize(substr)
		if _, ok := hashmap[substrNorm]; ok {
			return true
		}
	}

	return false

}

func normalize(runes []rune) [26]byte {
	res := [26]byte{}
	for _, r := range runes {
		ind := r - 'a'
		res[ind]++
	}
	return res
}
