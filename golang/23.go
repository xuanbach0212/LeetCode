//go:build ignore

package main

import (
	"fmt"
)

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
//
// Example 1:
//
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted linked list:
// 1->1->2->3->4->4->5->6
// Example 2:
//
// Input: lists = []
// Output: []
// Example 3:
//
// Input: lists = [[]]
// Output: []

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 || lists == nil {
		return nil
	}

	for len(lists) > 1 {
		var mergeLists []*ListNode

		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}
			mergeLists = append(mergeLists, mergeTwoLists(l1, l2))
		}

		lists = mergeLists
	}

	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	dummy := res

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
		dummy = dummy.Next
	}

	if list1 != nil {
		dummy.Next = list1
	}

	if list2 != nil {
		dummy.Next = list2
	}

	return res.Next
}

func main() {
	fmt.Println("mergeTwoLists is: ", mergeTwoLists(nil, nil))
}
