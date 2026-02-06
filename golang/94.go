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

// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
//
//
// Example 1:
//
// Input: root = [1,null,2,3]
//
// Output: [1,3,2]
//
// Explanation:
//()
//
//
// Example 2:
//
// Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
//
// Output: [4,2,6,5,7,1,3,9,8]
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
// Follow up: Recursive solution is trivial, could you do it iteratively?

func inorderTraversal(root *TreeNode) []int {
	// res := []int{}

	// var dfs func(node *TreeNode) *TreeNode
	// dfs = func(node *TreeNode) *TreeNode {
	// 	if node == nil {
	// 		return nil
	// 	}
	// 	dfs(node.Left)
	// 	res = append(res, node.Val)
	// 	dfs(node.Right)

	// 	return node
	// }
	// dfs(root)
	// return res

	// follow up
	res := []int{}
	stack := []*TreeNode{}
	curr := root
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
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
	result := inorderTraversal(root)

	fmt.Println("inorderTraversal is:", result)
}
