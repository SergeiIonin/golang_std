package main

// https://leetcode.com/problems/valid-palindrome/

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

import (
	"fmt"
	"strings"
)

func main() {
	in0 := "A man, a plan, a canal: Panama"
	res0 := isValidPalindrome(in0)
	fmt.Println(res0)

	in1 := "race a car"
	res1 := isValidPalindrome(in1)
	fmt.Println(res1)

	in2 := " "
	res2 := isValidPalindrome(in2)
	fmt.Println(res2)

	in3 := "0P"
	res3 := isValidPalindrome(in3)
	fmt.Println(res3) // false

	in4 := "aa"
	res4 := isValidPalindrome(in4)
	fmt.Println(res4) // true
}

// we can use "unicode" package (isLetter, is Digit, toLower methods)

// checks if a string is a valid palindrome

func isValidPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	// transform string to slice of runes
	runes := []rune(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphanumeric(runes[i]) {
			i++
			continue
		}
		if !isAlphanumeric(runes[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// isAlphanumeric checks if a rune is a letter or a digit

func isAlphanumeric(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}

// git command to revert changes in this file
// git checkout -- leetcode_tasks/easy/validPalindrome/main.go

// git command to rollback changes in this file
// git reset HEAD leetcode_tasks/easy/validPalindrome/main.go
