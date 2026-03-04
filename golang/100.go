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

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//
//
//
// Example 1:
//
//
// Input: p = [1,2,3], q = [1,2,3]
// Output: true
// Example 2:
//
//
// Input: p = [1,2], q = [1,null,2]
// Output: false
// Example 3:
//
//
// Input: p = [1,2,1], q = [1,1,2]
// Output: false

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}

	if (q == nil && p != nil) || (q != nil && p == nil) {
		return false
	}

	if q.Val != p.Val {
		return false
	}

	if !isSameTree(p.Right, q.Right) ||
		!isSameTree(p.Left, q.Left) {
		return false
	}

	return true
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(5), intPtr(1), intPtr(6), intPtr(2), intPtr(0), intPtr(8), nil, nil, intPtr(7), intPtr(4),
	}
	root := BuildTree(vals)
	result := isSameTree(root, root)

	fmt.Println(" isSameTree is:", result)
}
