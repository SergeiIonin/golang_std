package main

import "fmt"

// https://leetcode.com/problems/group-anagrams/

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

func main() {
	in0 := []string{"eat", "tea", "tan", "ate", "nat", "bat"} //Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	r0 := groupAnagrams(in0)
	fmt.Println(r0)

	in1 := []string{""}
	r1 := groupAnagrams(in1)
	fmt.Println(r1)

	in2 := []string{"a"}
	r2 := groupAnagrams(in2)
	fmt.Println(r2)

	in3 := []string{"ac", "c"} // [["c"],["ac"]]
	r3 := groupAnagrams(in3)
	fmt.Println(r3)
}

// it works but doesn't perform very well, event poorly than the solution w/ map, bc the overhead for
// array (frequency array) seems to be bigger
func groupAnagrams(strs []string) [][]string {

	anagrams := make([][]string, 0, len(strs))
	var mem []rune
	added := make(map[int]struct{})
	k := 0

	for i := 0; i < len(strs); i++ {
		_, ok := added[i]
		if ok {
			continue
		}
		anagrams = append(anagrams, []string{strs[i]})
		k = len(anagrams) - 1
		mem = memoizeStr(strs[i])
		for j := i + 1; j < len(strs); j++ {
			_, ok := added[j]
			if ok {
				continue
			}
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			m := copyMem(mem)
			if isAnagram(m, strs[j]) {
				added[j] = struct{}{}
				anagrams[k] = append(anagrams[k], strs[j])
			}
		}
	}

	return anagrams
}

func memoizeStr(str string) []rune {
	m := make([]rune, 122) // bc all letters are english lowercase
	for _, r := range str {
		m[r]++
	}
	return m
}

func copyMem(src []rune) []rune {
	dst := make([]rune, len(src))
	copy(dst, src)
	return dst
}

func isAnagram(m []rune, str string) bool {
	for _, r := range str {
		res := m[r]
		if res == 0 {
			return false
		}
		m[r]--
	}
	return true
}
