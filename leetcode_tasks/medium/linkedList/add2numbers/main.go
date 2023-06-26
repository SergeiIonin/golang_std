package main

import "fmt"

// https://leetcode.com/problems/add-two-numbers/

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
// and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

func main() {
	// ListNode with 2,4,3
	// list1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// // ListNode with 5,6,4
	// list2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	// res12 := addTwoNumbers(list1, list2)
	// printListNode(res12) // [7 0 8]

	// listNode with 9,9,9,9,9,9,9
	list3 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	//9,9,9,9
	list4 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	res34 := addTwoNumbers(list3, list4)
	printListNode(res34) // [8,9,9,9,0,0,0,1]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	v1 := l1.Val
	v2 := l2.Val
	sum := v1 + v2
	rem := 0
	if sum > 9 {
		sum = sum - 10
		rem = 1
	}
	res := &ListNode{sum, nil}
	cur := res

	l1 = l1.Next
	l2 = l2.Next
	var temp *ListNode

	for l1 != nil || l2 != nil {
		v1 = getValueOrZero(l1)
		v2 = getValueOrZero(l2)
		sum = v1 + v2 + rem
		if sum > 9 {
			sum = sum - 10
			rem = 1
		} else {
			rem = 0
		}
		temp = &ListNode{sum, nil}
		cur.Next = temp
		cur = cur.Next
		l1 = nextOrNil(l1)
		l2 = nextOrNil(l2)
	}

	if rem == 1 {
		cur.Next = &ListNode{1, nil}
	}

	return res
}

func getValueOrZero(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func nextOrNil(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
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
