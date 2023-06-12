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
	res0 := isPalindrome(in0)
	fmt.Println(res0)

	in1 := "race a car"
	res1 := isPalindrome(in1)
	fmt.Println(res1)

	in2 := " "
	res2 := isPalindrome(in2)
	fmt.Println(res2)

	in3 := "0P"
	res3 := isPalindrome(in3)
	fmt.Println(res3) // false

	in4 := "aa"
	res4 := isPalindrome(in4)
	fmt.Println(res4) // true
}

// we can use "unicode" package (isLetter, is Digit, toLower methods)
func isPalindrome(s string) bool {
	slc := strings.ToLower(s)
	sb := strings.Builder{}
	for _, r := range slc {
		if (r >= 97 && r <= 122) || (r >= 48 && r <= 57) {
			sb.WriteRune(r)
		}
	}
	runes := []rune(sb.String())
	if len(runes) == 0 {
		return true
	}
	i := 0
	j := len(runes) - 1

	for i < j {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}

	return true
}
