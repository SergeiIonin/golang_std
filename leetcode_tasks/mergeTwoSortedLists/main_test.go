package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {

	inputs := []TestInputAndResult{
		{
			list1: &ListNode{1, &ListNode{3, nil}},
			list2: &ListNode{2, &ListNode{4, nil}},
			res:   &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			name:  "[1 3] and [2 4]",
		},
		{
			list1: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			list2: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			res:   &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
			name:  "[1 3 5] and [2 4 6]",
		},
		{
			list1: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			list2: nil,
			res:   &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			name:  "[1 3 5] and nil",
		},
		{
			list1: nil,
			list2: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			res:   &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			name:  "nil [2 4 6]",
		},
		{
			list1: nil,
			list2: nil,
			res:   nil,
			name:  "2 nils",
		},
	}
	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			got := mergeTwoLists(input.list1, input.list2)
			if got != input.res {
				fmt.Println("test ", input.name)
				fmt.Println("got = ")
				printListNode(got)
				t.Error("test failed")
				//t.Errorf("got: %s, expected: %s", got, input.res)
			}
		})
	}

}

type TestInputAndResult struct {
	list1 *ListNode
	list2 *ListNode
	res   *ListNode
	name  string
}
