package main

import (
	"fmt"
)

func main() {
	s := "weallloveyou"
	maxVowels(s, 7)
}
// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length
/* Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.

Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.

Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.

Constraints:
    1 <= s.length <= 10^5
    s consists of lowercase English letters.
    1 <= k <= s.length
 */
func maxVowels(s string, k int) int {
    vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	vowelsMap := make(map[rune]bool)

	for _, v := range vowels {
		vowelsMap[v] = true
	}

	l := 0
	r := l+k

	runes := []rune(s)

	windowInit := runes[l:r]
	count := 0

	isVowel := func(r rune) bool {
		return vowelsMap[r]
	}

	for _, elem := range windowInit {
		if isVowel(elem) {
			count++
		}
	}
	
	countMax := count

	if count == k {
		return k
	}

	updateCount := func(l, r int) {
		fmt.Printf("window = %v \n", string(runes[l:r]))
		fmt.Printf("count before = %d \n", count)
		fmt.Printf("moving %s to %s \n", string(runes[l]), string(runes[l+1]))
		if k > 1 {
			if isVowel(runes[l]) {
				count--
			}
		} else {
			if isVowel(runes[l+1]) {
				count = 1
			}
		}
		fmt.Printf("count_L = %d \n", count)
		fmt.Printf("moving %s to %s \n", string(runes[r-1]), string(runes[r]))
		if k > 1 {
			if isVowel(runes[r]) {
				count++
			}
		}
		fmt.Printf("count after = %d \n", count)
		fmt.Printf("----------\n")
		if count > countMax {
			countMax = count
		}
	}

	for ; r < len(runes)-1; {
		fmt.Printf("l = %d, r = %d \n", l, r)
		updateCount(l, r)
		l++
		r++
		if count == k {
			return k
		}
	}
	fmt.Printf("countMax = %d \n", countMax)
	return countMax
}