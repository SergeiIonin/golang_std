package main

import "fmt"

func main() {
	// call removeNthFromEnd with example data
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 2
	res := removeNthFromEnd(list, n)
	printListNode(res)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}
	nodes := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	size := len(nodes)
	index := size - n
	if index == 0 {
		return head.Next
	}
	for i, node := range nodes {
		if i == index-1 {
			node.Next = node.Next.Next
		}
	}
	return head
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
