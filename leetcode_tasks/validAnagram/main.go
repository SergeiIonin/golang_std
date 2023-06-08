package main

// https://leetcode.com/problems/valid-anagram

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runes_s := []rune(s)
	runes_t := []rune(t)
	size := len(runes_s)

	hash := make(map[rune]int, size)

	for i := 0; i < size; i++ {
		hash[runes_s[i]]++
	}

	for i := 0; i < size; i++ {
		v, ok := hash[runes_t[i]]
		if ok && v > 0 {
			hash[runes_t[i]]--
		} else {
			return false
		}
	}

	return true
}

// most performant solutioin from leetcode
// func isAnagram(s string, t string) bool {
//     if len(s) != len(t) {
//         return false
//     }

//     mem := make([]int, 256)
//     for _, v := range s{
//         mem[v] += 1
//     }
//     for _, v := range t{
//         mem[v] -= 1
//         if mem[v] < 0{
//             return false
//         }
//     }
//     return true
// }
