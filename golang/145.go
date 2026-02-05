//go:build ignore

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func intPtr(v int) *int { return &v }

// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
//
//
// Example 1:
//
// Input: root = [1,null,2,3]
//
// Output: [3,2,1]
//
// Explanation:
//
//
//
// Example 2:
//
// Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// Output: [4,6,7,5,2,9,8,3,1]
//
// Explanation:
//
//
//
// Example 3:
//
// Input: root = []
//
// Output: []
//
// Example 4:
//
// Input: root = [1]
//
// Output: [1]
//
//
//
// Constraints:
//
// The number of the nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
// Follow up: Recursive solution is trivial, could you do it iteratively?

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	node := root
	var lastVisited *TreeNode
	for len(stack) > 0 || node != nil {
		for node != nil {
			// still go farest left
			stack = append(stack, node)
			node = node.Left
		}
		peekNode := stack[len(stack)-1]

		if peekNode.Right != nil && lastVisited != peekNode.Right { // do we need to go right
			node = peekNode.Right
		} else {
			res = append(res, peekNode.Val)
			stack = stack[:len(stack)-1]
			lastVisited = peekNode
		}

	}

	return res
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4]
	// p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7),
	}
	root := BuildTree(vals)
	result := postorderTraversal(root)

	fmt.Println("postorderTraversal is:", result)
}
