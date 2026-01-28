//go:build ignore

package main

import "fmt"

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

// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
//
//
//
// Example 1:
//
//
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
// Example 2:
//
// Input: root = [1]
// Output: [[1]]
// Example 3:
//
// Input: root = []
// Output: []

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		qLen := len(queue)
		level := []int{}
		for range qLen {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				level = append(level, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}

		if len(level) > 0 {
			res = append(res, level)
		}
	}

	return res
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(5), intPtr(1), intPtr(6), intPtr(2), intPtr(0), intPtr(8), nil, nil, intPtr(7), intPtr(4),
	}
	root := BuildTree(vals)
	result := levelOrder(root)

	fmt.Println("levelOrder is:", result)
}
