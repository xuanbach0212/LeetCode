//go:build ignore

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	biggest := 0
	for prev != nil {
		tmp := head.Val + prev.Val
		if tmp > biggest {
			biggest = tmp
		}
		head = head.Next
		prev = prev.Next
	}
	return biggest
}

func main() {
	node4 := &ListNode{Val: 4, Next: nil}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	head := &ListNode{Val: 1, Next: node2}

	result := pairSum(head)

	fmt.Println("Maximum twin sum:", result)
}
