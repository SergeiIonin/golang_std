package main

import "fmt"

func main() {
	in := []int{1,0,0,0,1,0,0}
	res := kidsWithCandies(in, 2)
	fmt.Printf("res = %v\n", res)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    excludedIndexes := make([]int, 0, len(candies))
    res := make([]bool, len(candies))
    max := candies[0]
    for i, elem := range candies {
        if elem > max {
            max = elem
        }
        if (elem + extraCandies) >= max {
            excludedIndexes = append(excludedIndexes, i)
        }
    }

    for _, j := range excludedIndexes {
        if (candies[j] + extraCandies) >= max {
            res[j] = true
        }
    }

    return res
}