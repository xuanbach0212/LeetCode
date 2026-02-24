//go:build ignore

package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 4 5
// nil<-1<-
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// prev = nil curr = 1    begin  curr = 1, tmp = 1, curr = 2, tmp.next = prev = nil,

func main() {
}
