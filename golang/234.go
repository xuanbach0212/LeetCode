//go:build ignore

package main

import (
	"fmt"
)

// 234. Palindrome Linked List
// Easy
// Topics
// premium lock icon
// Companies
// Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
//
//
//
// Example 1:
//
//
// Input: head = [1,2,2,1]
// Output: true
// Example 2:
//
//
// Input: head = [1,2]
// Output: false
//
//
// Constraints:
//
// The number of nodes in the list is in the range [1, 105].
// 0 <= Node.val <= 9
//
//
// Follow up: Could you do it in O(n) time and O(1) space?

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil { // which mean the len is odd, so we skip first slow
		slow = slow.Next
	}

	var prev *ListNode
	// reverse 2nd half
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// compare
	for prev != nil {
		if prev.Val != head.Val {
			return false
		}

		prev = prev.Next
		head = head.Next
	}

	return true
}

func main() {
	fmt.Println("isPalindrome is: ", isPalindrome(nil))
}
