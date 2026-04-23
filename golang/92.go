//go:build ignore

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.
//
//
//
// Example 1:
//
//
// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]
// Example 2:
//
// Input: head = [5], left = 1, right = 1
// Output: [5]
//
//
// Constraints:
//
// The number of nodes in the list is n.
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
// Follow up: Could you do it in one pass?

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	before := dummy

	for range left - 1 {
		before = before.Next
	}

	curr := before.Next
	var prevRev *ListNode
	for range right - left + 1 {
		next := curr.Next
		curr.Next = prevRev
		prevRev = curr
		curr = next
	}

	before.Next.Next = curr
	before.Next = prevRev

	return dummy.Next
}

func main() {
	fmt.Println("removeNthFromEnd")
}
