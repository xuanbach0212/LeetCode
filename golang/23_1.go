//go:build ignore

package main

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
		var mergeList []*ListNode
		for i := 0; i < len(lists); i += 2 {
			list1 := lists[i]
			var list2 *ListNode
			if i+1 < len(lists) {
				list2 = lists[i+1]
			}
			mergeList = append(mergeList, merge2Lists(list1, list2))
		}
		lists = mergeList
	}
	return lists[0]
}

func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
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
}
