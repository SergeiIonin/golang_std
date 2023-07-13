package main

import "sort"

// https://leetcode.com/problems/car-fleet/

// There are n cars going to the same destination along a one-lane road. The destination is target miles away.

// You are given two integer array position and speed, both of length n, where position[i] is the position of the ith car and speed[i] is the speed of the ith car (in miles per hour).

// A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper at the same speed. The faster car will slow down to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).

// A car fleet is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.

// If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.

// Return the number of car fleets that will arrive at the destination.

func main() {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}

	carFleet(target, position, speed)
}

// based on the most performant solution on leetcode
func carFleet(target int, position []int, speed []int) int {

	posWithSpeed := make([][2]int, len(position))
	for i := 0; i < len(position); i++ {
		posWithSpeed[i] = [2]int{position[i], speed[i]}
	}

	// sort by position in descending order
	sort.Slice(posWithSpeed, func(i, j int) bool {
		return posWithSpeed[i][0] > posWithSpeed[j][0]
	})

	stack := make([]float64, 0, len(position))

	for _, elem := range posWithSpeed {
		t := float64((target - elem[0]) / elem[1])
		if len(stack) == 0 || t > stack[len(stack)-1] {
			stack = append(stack, t)
		}
	}

	return len(stack)

}
