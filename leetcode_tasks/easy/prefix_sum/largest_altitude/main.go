package main

import "fmt"

func main() {
	in := []int{-5,1,5,0,-7}
	res := largestAltitude(in)
	fmt.Printf("res = %v\n", res)
}

/*
https://leetcode.com/problems/find-the-highest-altitude
There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. 
The biker starts his trip on point 0 with altitude equal 0.

You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n).
Return the highest altitude of a point.

Constraints:

n == gain.length
1 <= n <= 100
-100 <= gain[i] <= 100
*/
func largestAltitude(gain []int) int {
    max := 0
    cur := gain[0]
    if cur > max {
        max = cur
    }

    for i := 1; i < len(gain); i++ {
        cur += gain[i]
        if cur > max {
            max = cur
        }
    }

    return max
    
}