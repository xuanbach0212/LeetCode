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

// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
//
//
// Example 1:
//
// Input: root = [1,null,2,3]
//
// Output: [1,2,3]
//
// Explanation:
//
//
//
// Example 2:
//
// Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// Output: [1,2,4,5,6,7,3,8,9]
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

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
		return node
	}
	dfs(root)
	return res
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4]
	// p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7),
	}
	root := BuildTree(vals)
	result := preorderTraversal(root)

	fmt.Println("preorderTraversal is:", result)
}
