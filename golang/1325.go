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

// Given a binary tree root and an integer target, delete all the leaf nodes with value target.
//
// Note that once you delete a leaf node with value target, if its parent node becomes a leaf node and has the value target, it should also be deleted (you need to continue doing that until you cannot).
//
//
//
// Example 1:
//
//
//
// Input: root = [1,2,3,2,null,2,4], target = 2
// Output: [1,null,3,null,4]
// Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left).
// After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).
// Example 2:
//
//
//
// Input: root = [1,3,3,3,2], target = 3
// Output: [1,3,null,null,2]
// Example 3:
//
//
//
// Input: root = [1,2,null,2,null,2], target = 2
// Output: [1]
// Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
//
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 3000].
// 1 <= Node.val, target <= 1000

//        1
//      2   3
//     2   2 4
//
//
//       1
//     3   3
//    3 2
//
//
// if node val = target then del
// if node val = target and left, right is null then del
// dfs
//
// time: O(n)
// space: O(log n)
//
// [1,1,1]

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	curr := root
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)

		if node.Left == nil && node.Right == nil && node.Val == target {
			return nil
		}

		return node
	}
	return dfs(curr)
}

func main() {
	// root = [3,1,4,3,null,1,5]
	node6 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 4, Left: node5, Right: node6}
	node2 := &TreeNode{Val: 1, Left: node4, Right: nil}
	head := &TreeNode{Val: 3, Left: node2, Right: node3}
	result := removeLeafNodes(head, 1)

	fmt.Println("removeLeafNodes:", result)
}
