//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.
//
// Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night.
//
// Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.
//
//
//
// Example 1:
//
//
// Input: root = [3,2,3,null,3,null,1]
// Output: 7
// Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
// Example 2:
//
//
// Input: root = [3,4,5,1,3,null,1]
// Output: 9
// Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
//
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 104].
// 0 <= Node.val <= 104
//
// [4,1,null,2,null,3]

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		withRLeft, withoutRLeft := dfs(node.Left)
		withRRight, withoutRRight := dfs(node.Right)
		return node.Val + withoutRLeft + withoutRRight, max(withRLeft, withoutRLeft) + max(withRRight, withoutRRight)
	}
	withR, withoutR := dfs(root)
	return max(withR, withoutR)
}

// TC: O(n)
// SC: O(log n)

//           3
//        4     5
//      1  3      1

// node = 1 => return 1, 0
// node = 3 => return 3, 0
// node = 4 => return 4, 4
// node = 1 => return 1, 0
// node = 5 => return 5, 1
// node = 3 => return 8, 9
//
//
//          4
//       1
//    2
//  3
//
// node = 3 => return 3, 0
// node = 2 => return 2, 3 + 0
// node = 1 => return 3 + 1 = 4, 5
// node = 4 => return

func main() {
	// root = [3,1,4,3,null,1,5]
	node6 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 4, Left: node5, Right: node6}
	node2 := &TreeNode{Val: 1, Left: node4, Right: nil}
	head := &TreeNode{Val: 3, Left: node2, Right: node3}
	result := rob(head)

	fmt.Println("rob:", result)
}
