package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	res := isSubsequence(s, t)
	fmt.Printf("res = %v\n", res)
}

/*
https://leetcode.com/problems/is-subsequence
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

Constraints:
    0 <= s.length <= 100
    0 <= t.length <= 10^4
    s and t consist only of lowercase English letters.
*/
func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	i := 0
	j := 0
	rs := []rune(s)
	ts := []rune(t)
	for ; j < len(t); {
		if rs[i] == ts[j] {
			i++
		}
		j++
		if i == len(s) {
			return true
		}
	}
	return false
}
