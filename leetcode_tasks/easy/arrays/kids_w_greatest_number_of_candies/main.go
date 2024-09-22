package main

import "fmt"

func main() {
	in := []int{2,3,5,1,3}
	res := kidsWithCandies(in, 3)
	fmt.Printf("res = %v\n", res) // [true, true, true, false, true]
}

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies
/* Constraints:

    n == candies.length
    2 <= n <= 100
    1 <= candies[i] <= 100
    1 <= extraCandies <= 50
 */
func kidsWithCandies(candies []int, extraCandies int) []bool {
    // indexes i for which candies[i]+extraCandies >= max, where max is the current max at index i, so we'll need
    // to walk these indexes once more after determination of the absolute max
    excludedIndexes := make([]int, 0, len(candies))
    res := make([]bool, len(candies))
    max := candies[0]
    excludedIndexes = append(excludedIndexes, 0)

    for i, elem := range candies[1:] {
        if elem > max {
            max = elem
        }
        if (elem + extraCandies) >= max {
            excludedIndexes = append(excludedIndexes, i+1)
        }
    }

    for _, j := range excludedIndexes {
        if (candies[j] + extraCandies) >= max {
            res[j] = true
        }
    }

    return res
}