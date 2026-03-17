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

// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
//
//
//
// Example 1:
//
// Input: root = [1,2,3,null,5,null,4]
//
// Output: [1,3,4]
//
// Explanation:
//
//
//
// Example 2:
//
// Input: root = [1,2,3,4,null,null,null,5]
//
// Output: [1,3,4,5]
//
// Explanation:
//
//
//
// Example 3:
//
// Input: root = [1,null,3]
//
// Output: [1,3]
//
// Example 4:
//
// Input: root = []
//
// Output: []
//
//
//
// Constraints:
//
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100

func rightSideView(root *TreeNode) []int {
	// use iterative

	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	res := []int{root.Val}

	for len(queue) > 0 {
		qLen := len(queue)
		for range qLen {
			curr := queue[0]
			queue = queue[1:]

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		if len(queue) > 0 {
			res = append(res, queue[len(queue)-1].Val)
		}
	}
	return res
}

func main() {
	// root = [1,2,3,null,5,null,4]
	vals := []*int{
		intPtr(1), intPtr(2), intPtr(3), nil, intPtr(5), nil, intPtr(4),
	}
	root := BuildTree(vals)
	result := rightSideView(root)

	for _, v := range result {
		fmt.Println(" rightSideView is:", v)
	}
}
