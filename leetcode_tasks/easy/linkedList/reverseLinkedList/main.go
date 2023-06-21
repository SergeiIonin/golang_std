package main

import "fmt"

func main() {

	// a slice of 2 TestDatas

	testData := []TestData{
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			expected: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
			name:     "test1",
		},
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
			name:     "test2",
		},
	}

	testReverseList(testData)
}

// Given the head of a singly linked list, reverse the list, and return the reversed list.

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next // save the next node
		curr.Next = prev  // reverse the current node
		prev = curr       // move prev to curr
		curr = next       // move curr to next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// let's write a test for this task

// test struct which contains input data, expected result and name of the test

type TestData struct {
	input    *ListNode
	expected *ListNode
	name     string
}

// write a function which accepts a slice of TestData and check each test

func testReverseList(testData []TestData) {
	for _, test := range testData {
		result := reverseList(test.input)
		if !isSame(result, test.expected) {
			fmt.Println("test", test.name, "failed")
			fmt.Println("expected", test.expected, "got", result)
			return
		}
	}
	fmt.Println("all tests passed")
}

// write isSame function which checks if two linked lists are the same

func isSame(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 != nil && l2 == nil || l1 == nil && l2 != nil {
		return false
	}
	if l1.Val != l2.Val {
		return false
	}
	return isSame(l1.Next, l2.Next)
}
