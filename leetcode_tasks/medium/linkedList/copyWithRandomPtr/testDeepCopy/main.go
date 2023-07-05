package main

import "fmt"

func main() {
	initial := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}

	copy := deepCopyListNode(initial)

	printListNode(copy)
	copy.Next = &ListNode{5, nil}

	printListNode(copy)
	printListNode(initial) // initial doesn't change
}

func deepCopyListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return &ListNode{head.Val, deepCopyListNode(head.Next)}
}

func deepCopy(head *ListNode) *ListNode {

	s := size(head)

	buf := make([]*ListNode, s)
	buf[0] = &ListNode{head.Val, nil}
	index := 1

	cur := head.Next
	for cur != nil {
		v := &ListNode{cur.Val, nil}
		buf[index] = v
		buf[index-1].Next = v
		index++
		cur = cur.Next
	}

	return buf[0]

}

func size(head *ListNode) int {
	if head == nil {
		return 0
	}
	return 1 + size(head.Next)
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

type ListNode struct {
	Val  int
	Next *ListNode
}
