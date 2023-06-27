package main

import "fmt"

//https://leetcode.com/problems/copy-list-with-random-pointer/

// A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

// Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of
// its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers
// in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

// For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list,
// x.random --> y.

// Return the head of the copied linked list.

// The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:
// val: an integer representing Node.val
// random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
// Your code will only be given the head of the original linked list.

func main() {

	// Node with elements 7,13,11,10,1 where 7's random points to nil, 13's random points to 7, 11's random points to 10, 10's random points to 11, 1's random points to 7

	n0 := &Node{7, nil, nil}
	n1 := &Node{13, nil, n0}
	n2 := &Node{11, nil, nil}
	n3 := &Node{10, nil, n2}
	n4 := &Node{1, nil, n0}

	n2.Random = n4
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	copy := copyRandomList(n0)

	printListNode(copy)

}

// it works but pretty slow and consumes a lot of memory
func copyRandomList(head *Node) *Node {
	hashmap := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		hashmap[cur] = &Node{cur.Val, nil, nil}
		cur = cur.Next
	}

	cur = head

	for cur != nil {
		copy := hashmap[cur]
		copy.Next = hashmap[cur.Next]
		copy.Random = hashmap[cur.Random]
		cur = cur.Next
	}

	return hashmap[head]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func printListNode(ln *Node) {
	var accum [][2]int
	if ln == nil {
		fmt.Println("[]")
	}
	for {
		rand := ln.Random
		randVal := 100000
		if rand != nil {
			randVal = rand.Val
		}
		accum = append(accum, [2]int{ln.Val, randVal})
		ln = ln.Next
		if ln == nil {
			break
		}
	}
	fmt.Println(accum)
}
