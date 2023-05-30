package main

import "fmt"

func main() {
	s1 := "MMMCMXCIX"
	s2 := "DCCCLXXXVIII"
	s3 := "LVIII"
	s4 := "MDLXX"
	s5 := "MMM"

	int1 := romanToInt(s1)
	int2 := romanToInt(s2)
	int3 := romanToInt(s3)
	int4 := romanToInt(s4)
	int5 := romanToInt(s5)

	fmt.Printf("%s => %d \n", s1, int1)
	fmt.Printf("%s => %d \n", s2, int2)
	fmt.Printf("%s => %d \n", s3, int3)
	fmt.Printf("%s => %d \n", s4, int4)
	fmt.Printf("%s => %d \n", s5, int5)
}

// [1, 3999]
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// MMM CM XC IX = 3999
// DCCC LXXX VIII = 888

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// if starts from I..V 		  		1..10
//		if starts from IV | IX 		4 | 9

// if starts from X 				>= 10
//		if starts from XL | XC		40..90

// if starts from C			        100..900
// if starts from CD || CM  		400..900

// if starts from M		  			1000...3999

func romanToInt(s string) int {
	runes := []rune(s)
	sizeRunes := int8(len(s))

	accumArray := make([]rune, 0, 4)

	accum := 0
	var cur int8 = 3

	var currentRune rune

	var index int8 = 0
	for ; ; index++ {
		if index == sizeRunes {
			break
		}
		currentRune = runes[index]
		switch currentRune {
		case 'M':
			accumArray = append(accumArray, 'M')

		case 'D':
			if cur == 3 {
				cur = 2
				accum += len(accumArray) * 1000
				accumArray = make([]rune, 0, 4)
			}
			accumArray = append(accumArray, 'D')
		case 'C':
			if cur == 3 {
				cur = 2
				accum += len(accumArray) * 1000
				accumArray = make([]rune, 0, 4)
			}
			accumArray = append(accumArray, 'C')

		case 'L':
			if cur > 1 {
				cur = 1
				accum += convert1000(accumArray)
				accum += convert(accumArray, 'C', 'D', 'M', 100)
				accumArray = make([]rune, 0, 4)
			}
			accumArray = append(accumArray, 'L')
		case 'X':
			if cur > 1 {
				cur = 1
				accum += convert1000(accumArray)
				accum += convert(accumArray, 'C', 'D', 'M', 100)
				accumArray = make([]rune, 0, 4)
			}
			accumArray = append(accumArray, 'X')

		case 'V':
			if cur >= 1 {
				cur = 0
				accum += convert1000(accumArray)
				accum += convert(accumArray, 'C', 'D', 'M', 100)
				accum += convert(accumArray, 'X', 'L', 'C', 10)
				accumArray = make([]rune, 0, 4)
			}
			accumArray = append(accumArray, 'V')
		case 'I':
			if cur >= 1 {
				cur = 0
				accum += convert1000(accumArray)
				accum += convert(accumArray, 'C', 'D', 'M', 100)
				accum += convert(accumArray, 'X', 'L', 'C', 10)
				accumArray = make([]rune, 0, 4)
			}
			accumArray = append(accumArray, 'I')
		}
	}

	addition := convert(accumArray, 'I', 'V', 'X', 1)
	if addition > 0 {
		return accum + addition
	}
	addition = convert(accumArray, 'X', 'L', 'C', 10)
	if addition > 0 {
		return accum + addition
	}
	addition = convert(accumArray, 'C', 'D', 'M', 100)
	if addition > 0 {
		return accum + addition
	}
	addition = convert1000(accumArray)
	return accum + addition
}

func convert1000(runes []rune) (res int) {
	size := len(runes)
	count := 0
	for i := 0; i < size; i++ {
		if runes[i] == 'M' {
			count++
		} else {
			break
		}
	}
	return count * 1000
}

func convert(runes []rune, l rune, m rune, r rune, base int) (res int) {
	size := len(runes)
	if len(runes) == 0 {
		return
	}
	if runes[0] == l {
		if size == 1 {
			res = 1 * base
			return
		}
		if runes[1] == m {
			res = 4 * base
			return
		}
		if runes[1] == r {
			res = 9 * base
			return
		}
		res = len(runes) * base
		return
	}

	if runes[0] == m {
		res = (5 + (len(runes) - 1)) * base
		return
	}
	return
}
