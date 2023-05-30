package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := "MMMCMXCIX"    // 3999
	s2 := "DCCCLXXXVIII" // 888
	s3 := "LVIII"        // 58
	s4 := "MDLXX"        // 1570
	s5 := "MMM"          // 3000
	start := time.Now()
	int1 := romanToInt(s1)
	int2 := romanToInt(s2)
	int3 := romanToInt(s3)
	int4 := romanToInt(s4)
	int5 := romanToInt(s5)
	timeElapsed := time.Since(start)

	fmt.Printf("%s => %d \n", s1, int1)
	fmt.Printf("%s => %d \n", s2, int2)
	fmt.Printf("%s => %d \n", s3, int3)
	fmt.Printf("%s => %d \n", s4, int4)
	fmt.Printf("%s => %d \n", s5, int5)
	fmt.Printf("Time elapsed = %s", timeElapsed)
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
	for {
		if index == sizeRunes {
			break
		}
		currentRune = runes[index]
		if cur == 3 {
			if currentRune != 'M' {
				accum += len(accumArray) * 1000
				accumArray = cleanArray(accumArray)
				if currentRune == 'D' || currentRune == 'C' {
					cur = 2
				} else if currentRune == 'L' || currentRune == 'X' {
					cur = 1
				} else {
					cur = 0
				}
			}
			index++
			accumArray = append(accumArray, currentRune)
			continue
		}
		if cur == 2 {
			if currentRune != 'C' && currentRune != 'D' && currentRune != 'M' {
				accum += convert(accumArray, 'C', 'D', 'M', 100)
				accumArray = cleanArray(accumArray)
				if currentRune == 'X' || currentRune == 'L' {
					cur = 1
				} else {
					cur = 0
				}
			}
			index++
			accumArray = append(accumArray, currentRune)
			continue
		}
		if cur == 1 {
			if currentRune != 'X' && currentRune != 'L' && currentRune != 'C' {
				accum += convert(accumArray, 'X', 'L', 'C', 10)
				accumArray = cleanArray(accumArray)
				cur = 0
			}
			index++
			accumArray = append(accumArray, currentRune)
			continue
		}
		index++
		accumArray = append(accumArray, currentRune)
	}

	var res int
	switch cur {
	case 3:
		res = accum + len(accumArray)*1000
	case 2:
		res = accum + convert(accumArray, 'C', 'D', 'M', 100)
	case 1:
		res = accum + convert(accumArray, 'X', 'L', 'C', 10)
	case 0:
		res = accum + convert(accumArray, 'I', 'V', 'X', 1)
	}
	return res
}

func cleanArray(arr []rune) []rune {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

func convert(runes []rune, l rune, m rune, r rune, base int) (res int) {
	size := len(runes)
	if len(runes) == 0 {
		return
	}
	if runes[0] == l {
		if size == 1 {
			res = base
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
