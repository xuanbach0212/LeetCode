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

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
//
//
// Example 1:
//
//
// Input: root = [1,2,2,3,4,4,3]
// Output: true
// Example 2:
//
//
// Input: root = [1,2,2,null,3,null,3]
// Output: false
//
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
//
//
// Follow up: Could you solve it both recursively and iteratively?

func isSymmetric(root *TreeNode) bool {
	var dfs func(left *TreeNode, right *TreeNode) bool
	dfs = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}

func main() {
	// root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
	vals := []*int{
		intPtr(3), intPtr(5), intPtr(1), intPtr(6), intPtr(2), intPtr(0), intPtr(8), nil, nil, intPtr(7), intPtr(4),
	}
	root := BuildTree(vals)
	result := isSymmetric(root)

	fmt.Println("isSymmetric is:", result)
}
