package main

import "fmt"

func main() {
	list1 := &ListNode{1, &ListNode{3, nil}}
	list2 := &ListNode{2, &ListNode{4, nil}}

	m := mergeTwoLists(list1, list2)

	printListNode(m)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	///
	if list1 == nil {
		if list2 == nil {
			return nil
		} else {
			return list2
		}
	} else if list2 == nil {
		return list1
	}
	///
	val1_0 := list1.Val
	val2_0 := list2.Val
	var merged *ListNode
	if val1_0 < val2_0 {
		cur := val1_0
		merged = &ListNode{
			cur,
			nil,
		}
		list1 = list1.Next
	} else {
		cur := val2_0
		merged = &ListNode{
			cur,
			nil,
		}
		list2 = list2.Next
	}
	///
	var val1 int
	var val2 int
	next := merged

	for list1 != nil || list2 != nil { // list1.Next != nil && list2.Next != nil
		if list1 == nil {
			val2 = list2.Val
			cur := val2
			upd := &ListNode{
				cur,
				nil,
			}
			updateNext(next, upd)
			list2 = list2.Next
		} else if list2 == nil {
			val1 = list1.Val
			cur := val1
			upd := &ListNode{
				cur,
				nil,
			}
			updateNext(next, upd)
			list1 = list1.Next
		} else {
			val1 := list1.Val
			val2 := list2.Val

			if val1 < val2 {
				cur := val1
				upd := &ListNode{
					cur,
					nil,
				}
				updateNext(next, upd)
				list1 = list1.Next
			} else {
				cur := val2
				upd := &ListNode{
					cur,
					nil,
				}
				updateNext(next, upd)
				list2 = list2.Next
			}
		}
		next = next.Next
	}
	///
	return merged
}

func updateNext(next *ListNode, upd *ListNode) *ListNode {
	next.Next = upd
	return next
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
