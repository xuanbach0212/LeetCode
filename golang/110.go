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

// Given a binary tree, determine if it is height-balanced.
//
//
//
// Example 1:
//
//
// Input: root = [3,9,20,null,null,15,7]
// Output: true
// Example 2:
//
//
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false
// Example 3:
//
// Input: root = []
// Output: true

func isBalanced(root *TreeNode) bool {
	isValid := true
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		maxL := dfs(node.Left)
		maxR := dfs(node.Right)
		if maxL-maxR < -1 || maxL-maxR > 1 {
			isValid = false
		}
		return max(maxL, maxR) + 1
	}
	dfs(root)
	return isValid
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4]
	// p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7),
	}
	root := BuildTree(vals)
	result := isBalanced(root)

	fmt.Println("isBalanced is:", result)
}
