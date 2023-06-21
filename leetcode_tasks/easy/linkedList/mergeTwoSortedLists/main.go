package main

import "fmt"

func main() {
	list1 := &ListNode{1, &ListNode{3, nil}}
	list2 := &ListNode{2, &ListNode{4, nil}}

	m := mergeTwoLists(list1, list2)

	printListNode(m)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
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
