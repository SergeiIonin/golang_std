package main

import "fmt"

// https://leetcode.com/problems/group-anagrams/

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

func main() {
	in0 := []string{"eat", "tea", "tan", "ate", "nat", "bat"} //Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	r0 := groupAnagrams_(in0)
	fmt.Println(r0)

	in1 := []string{""}
	r1 := groupAnagrams_(in1)
	fmt.Println(r1)

	in2 := []string{"a"}
	r2 := groupAnagrams_(in2)
	fmt.Println(r2)

	in3 := []string{"ac", "c"} // [["c"],["ac"]]
	r3 := groupAnagrams_(in3)
	fmt.Println(r3)
}

// it works but doesn't perform very well
func groupAnagrams(strs []string) [][]string {

	anagrams := make([][]string, 0, len(strs))
	var hash map[rune]int
	added := make(map[int]struct{})
	k := 0

	for i := 0; i < len(strs); i++ {
		_, ok := added[i]
		if ok {
			continue
		}
		anagrams = append(anagrams, []string{strs[i]})
		k = len(anagrams) - 1
		hash = hashifyStr(strs[i])
		for j := i + 1; j < len(strs); j++ {
			_, ok := added[j]
			if ok {
				continue
			}
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			h := copyMap(hash)
			if isAnagram(h, strs[j]) {
				added[j] = struct{}{}
				anagrams[k] = append(anagrams[k], strs[j])
			}
		}
	}

	return anagrams
}

func hashifyStr(str string) map[rune]int {
	h := make(map[rune]int, len(str))
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		h[runes[i]]++
	}
	return h
}

func copyMap(src map[rune]int) map[rune]int {
	dst := make(map[rune]int, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func isAnagram(h map[rune]int, str string) bool {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		h[runes[i]]--
		if h[runes[i]] < 0 {
			return false
		}
	}
	return true
}

// this is based on the most performant solution on leetcode

func groupAnagrams_(strs []string) [][]string {
	hash := make(map[[26]byte][]string, len(strs))
	for _, s := range strs {
		norm := normalize(s)
		hash[norm] = append(hash[norm], s)
	}
	res := make([][]string, 0, len(strs))
	for _, strs := range hash {
		res = append(res, strs)
	}
	return res
}

func normalize(str string) [26]byte {
	res := [26]byte{}
	for _, r := range str {
		j := r - 'a'
		res[j]++
	}
	return res
}
