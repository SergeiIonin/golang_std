package main

import (
	"math"
)

func main() {

}

type MinStack struct {
	underlying []int
	min        int
}

func Constructor() MinStack {
	return MinStack{
		underlying: make([]int, 0, 10),
		min:        math.MaxInt,
	}
}

func (this *MinStack) Push(val int) {
	if val < this.min {
		this.min = val
	}
	this.underlying = append(this.underlying, val)
}

func (this *MinStack) Pop() {
	last := this.underlying[len(this.underlying)-1]
	this.underlying = this.underlying[:len(this.underlying)-1]
	if last == this.min {
		if len(this.underlying) > 0 {
			min := this.underlying[0]
			for i := 0; i < len(this.underlying); i++ {
				if this.underlying[i] < min {
					min = this.underlying[i]
				}
			}
			this.min = min
		} else {
			this.min = math.MaxInt
		}
	}
}

func (this *MinStack) Top() int {
	return this.underlying[len(this.underlying)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

// Most performant leetcode solution
// NB *Node are used, min is always fetched at O(1)
// type Node struct {
//     Val int
//     Prev *Node
//     Min int
// }

// type MinStack struct {
//     head *Node
// }

// func Constructor() MinStack {
//     return MinStack{
//         head: &Node{
//             Min: math.MaxInt32,
//         },
//     }
// }

// func (this *MinStack) Push(val int)  {
//     min := Min(this.head.Min, val)
//     this.head = &Node{
//         Val: val,
//         Prev: this.head,
//         Min: min,
//     }
// }

// func (this *MinStack) Pop()  {
//     this.head = this.head.Prev
// }

// func (this *MinStack) Top() int {
//     return this.head.Val
// }

// func (this *MinStack) GetMin() int {
//     return this.head.Min
// }

// func Min(a,b int) int {
//     if a<b { return a }
//     return b
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
