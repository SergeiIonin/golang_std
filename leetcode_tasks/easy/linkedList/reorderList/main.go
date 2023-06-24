package main

import "fmt"

// https://leetcode.com/problems/reorder-list/

// You are given the head of a singly linked-list. The list can be represented as:

// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:

// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.

func main() {
	// ListNode of elements 1,3,5,7,9
	list1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, nil}}}}}

	reorderList(list1)
	printListNode(list1)
}

func reorderList(head *ListNode) {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	reversed := reverse(slow.Next)
	slow.Next = nil

	curr := head

	for curr != nil && reversed != nil {
		next := curr.Next
		revNext := reversed.Next
		curr.Next = reversed
		reversed.Next = next
		curr = next
		reversed = revNext
	}
}

func reverse(node *ListNode) *ListNode {
	var prev, curr *ListNode = nil, node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(ln *ListNode) {
	var accum []int
	if ln == nil {
		fmt.Println("[]")
	}
	for {
		accum = append(accum, ln.Val)
		ln = ln.Next
		if ln == nil {
			break
		}
	}
	fmt.Println(accum)
}
