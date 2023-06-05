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
		if list2 == nil {
			return nil
		} else {
			return list2
		}
	} else if list2 == nil {
		return list1
	}
	val1_0 := list1.Val
	val2_0 := list2.Val
	var merged *ListNode
	if val1_0 < val2_0 {
		merged = &ListNode{
			val1_0,
			nil,
		}
		list1 = list1.Next
	} else {
		merged = &ListNode{
			val2_0,
			nil,
		}
		list2 = list2.Next
	}

	var val1 int
	var val2 int
	next := merged

	for list1 != nil || list2 != nil {
		if list1 == nil {
			val2 = list2.Val
			next.Next = &ListNode{
				val2,
				nil,
			}
			list2 = list2.Next
		} else if list2 == nil {
			val1 = list1.Val
			next.Next = &ListNode{
				val1,
				nil,
			}
			list1 = list1.Next
		} else {
			val1 := list1.Val
			val2 := list2.Val

			if val1 < val2 {
				next.Next = &ListNode{
					val1,
					nil,
				}
				list1 = list1.Next
			} else {
				next.Next = &ListNode{
					val2,
					nil,
				}
				list2 = list2.Next
			}
		}
		next = next.Next
	}
	return merged
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
