//go:build ignore

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a linked list, remove the nth node from the end of the list and return its head.
//
//
//
// Example 1:
//
//
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:
//
// Input: head = [1], n = 1
// Output: []
// Example 3:
//
// Input: head = [1,2], n = 1
// Output: [1]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	fmt.Println("removeNthFromEnd")
}
