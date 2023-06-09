package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/top-k-frequent-elements/

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

func main() {
	in0 := []int{1, 1, 1, 2, 2, 3}
	r0 := topKFrequent(in0, 2)
	fmt.Println(r0)
}

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}

	h := &MaxHeap{}

	heap.Init(h)

	res := make([]int, 0, k)

	for n, f := range hash {
		heap.Push(h, NF{n, f})
		if h.Len() > (len(hash) - k) {
			res = append(res, heap.Pop(h).(NF).num)
		}
	}

	return res

}

type NF struct {
	num  int
	freq int
}

type MaxHeap []NF

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(NF)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// fastest solution on leetcode

// func topKFrequent(nums []int, k int) []int {
//     numsMap := make(map[int]int)
//     result := []int{}

//     for _, val := range nums {
//         numsMap[val]++
//     }

//     // invert the previous map so we're keying by the frequency of times the number shows up
//     // need to use a list of ints as the value since a number could have the same frequency
//     freqMap := make(map[int][]int)
//     for key, freq := range numsMap {
//         freqMap[freq] = append(freqMap[freq], key)
//     }
// we can use the property that frequency values are bounded within [0,len(nums)]
// 	for i := len(nums); len(result) != k; i-- {
// 		for _, n := range freqMap[i] {
// 			if len(result) != k {
// 				result = append(result, n)
// 			}
// 		}
// 	}

//     return result
// }
