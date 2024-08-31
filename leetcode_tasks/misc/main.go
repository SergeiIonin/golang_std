package main

import (
	"fmt"
	"sort"
)


func main() {
	in := []int{1,3,5,8,9,2,6,7,6,8,9}

	res := minDifference(in, 7)

	fmt.Printf("res = %d \n", res)
}

func minDifference(earnings []int, k int) int {

	// The key of this map is a pair of integers [low, up], where low is the lowest possible bound for earnings[i] (either earnings[i] or earnings[i]-k)
	// and up is the upper possible bound for earnings[i] (which is always earnings[i]+k).		
	// The value of this map is [][2]int, this pair is (elemUp, elemUp) if the "bar" for earnings[i] is all below the key-"bar" (for earning[j]):
	// (k = 2)
	//		 _ up
	//		| |	
	//		|_|	
	//		| |	
	//      |_| low
	//	
	//	_  elemUp
	// | | 
	// |_| 
	// | |  
	// |_| elemLow
	//
	// If "bars" are overlapping but the key-"bar" is above the earnings[i]-"bar", then the pair is (elemLow, elemUp):
	//		       _ up
	//		      | |	
	//			  |_|	
	//	_ elemUp  | |	
	// | |  	  |_| low  
	// |_| 
	// | | 
	// |_| elemLow
	//
	// Hence, values in the pair are the biggest possible lower bound and the smallest possible upper bound for earnings[i].
	// And it means that if key-"bar" presents the highest bar, the pairs are the new possible values of earnings[i] given 
	// that the new value of key is either decreased or increased by k. If the lowest/highest value in the pair in unreachable, put -1,
	// if both are unreachable, then this key is invalid.
	// Ultimately this map should contain just one element, if the key contains an element (-1, -1), then this key is invalid.
	mapping := make(map[[2]int][][2]int)

	// create new earnings by deleting all repeating elements
	uniqueEarnings := make(map[int]bool)
	for _, elem := range earnings {
		uniqueEarnings[elem] = true
	}
	newEarnings := make([]int, 0, len(uniqueEarnings))
	for elem, _ := range uniqueEarnings {
		newEarnings = append(newEarnings, elem)
	}
	earnings = newEarnings

	if len(earnings) == 0 {
		return 0
	}

	low := 0
	up := 0
	for _, elem := range earnings {
		if elem < k {
			low = elem
		} else {
			low = elem-k
		}
		up = elem + k
		key := [2]int{low, up}
		mapping[key] = make([][2]int, 0)
	}

	sort.Ints(earnings)

	addToMapping := func(elem int) {
		for key, _ := range mapping {
			res := [2]int{-1, -1}
			low = key[0]
			up  = key[1]
			
			var elemLow int
			if elem < k {
				elemLow = elem
			} else {
				elemLow = elem - k
			}
			elemUp  := elem + k

			if elemUp <= low {
				res[0] = elemUp
				res[1] = elemUp
			} else {
				if elemLow <= low {
					res[0] = elemLow
					if elemUp <= up {
						res[1] = elemUp
					} else {
						res[1] = -1
					}
				} else if elemUp <= up {
					res[1] = elemUp	
				}
			}

			if res[0] == -1 && res[1] == -1 {
				delete(mapping, key)
			} else {
				mapping[key] = append(mapping[key], res)
			}
		}
	}
    
	for _, elem := range earnings {
        addToMapping(elem)
    }


	fmt.Printf("mapping = %v \n", mapping)

	lenMapping := len(mapping)
	deltas := make([]int, 0, lenMapping)

	for k, pairs := range mapping {
		minLow := pairs[0][0]
		maxLow := k[0]

		minUp := pairs[0][1]
		maxUp := k[1]
		for _, pair := range pairs {
			if pair[0] == k[0] && pair[1] == k[1] {
				continue
			}
			if pair[0] < minLow {
				minLow = pair[0]
			}
			if pair[0] > maxLow {
				maxLow = pair[0]
			}

			if pair[1] < minUp {
				minUp = pair[1]
			}
		}
		fmt.Printf("k = %v, minLow = %v, maxLow = %v, minUp = %v, maxUp = %v \n", k, minLow, maxLow, minUp, maxUp)
		var deltaLow int
		if minLow < 0 {
			deltaLow = -1
		} else {
			deltaLow = maxLow - minLow
		}
		deltaUp := maxUp - minUp
		fmt.Print("deltaLow = ", deltaLow, "; deltaUp = ", deltaUp, "\n")
		var delta int
		if deltaLow == -1 {
			delta = deltaUp
		} else {
			delta = min(deltaLow, deltaUp)
		}

		deltas = append(deltas, delta)
	}

	fmt.Printf("deltas = %v \n", deltas)
	min := deltas[0]

	for _, elem := range deltas {
		if elem < min {
			min = elem
		}
	}

    return min
}
