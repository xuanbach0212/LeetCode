//go:build ignore

package main

import (
	"fmt"
)

// You are given the head of a singly linked-list. The list can be represented as:
//
// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:
//
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.
//
//
//
// Example 1:
//
//
// Input: head = [1,2,3,4]
// Output: [1,4,2,3]
// Example 2:
//
//
// Input: head = [1,2,3,4,5]
// Output: [1,5,2,4,3]

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	first := head
	var prev *ListNode = nil
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	second := prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func main() {
	fmt.Println("reorderList")
}
