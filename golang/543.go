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

// Given the root of a binary tree, return the length of the diameter of the tree.
//
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
//
// The length of a path between two nodes is represented by the number of edges between them.
//
//
//
// Example 1:
//
//
// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
// Example 2:
//
// Input: root = [1,2]
// Output: 1

func diameterOfBinaryTree(root *TreeNode) int {
	// use dfs
	// go most left cal the max
	// then go right cal the max
	// cal the max when not split and split
	// return father node

	res := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		maxL := dfs(node.Left)
		maxR := dfs(node.Right)

		res = max(maxL+maxR, res)

		return 1 + max(maxL, maxR)
	}
	dfs(root)

	return res
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4]
	// p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(5), intPtr(1), intPtr(6), intPtr(2), intPtr(0), intPtr(8), nil, nil, intPtr(7), intPtr(4),
	}
	root := BuildTree(vals)
	result := diameterOfBinaryTree(root)

	fmt.Println("diameterOfBinaryTree is:", result)
}
