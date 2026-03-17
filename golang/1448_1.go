//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.
//
// Return the number of good nodes in the binary tree.
//
//
//
// Example 1:
//
//
//
// Input: root = [3,1,4,3,null,1,5]
// Output: 4
// Explanation: Nodes in blue are good.
// Root Node (3) is always a good node.
// Node 4 -> (3,4) is the maximum value in the path starting from the root.
// Node 5 -> (3,4,5) is the maximum value in the path
// Node 3 -> (3,1,3) is the maximum value in the path.
// Example 2:
//
//
//
// Input: root = [3,3,null,4,2]
// Output: 3
// Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
// Example 3:
//
// Input: root = [1]
// Output: 1
// Explanation: Root is considered as good.
//
//
// Constraints:
//
// The number of nodes in the binary tree is in the range [1, 10^5].
// Each node's value is between [-10^4, 10^4].

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxV int) int
	dfs = func(node *TreeNode, maxV int) int {
		if node == nil {
			return 0
		}
		count := 0
		if node.Val >= maxV {
			maxV = node.Val
			count += 1
		}
		count += dfs(node.Left, maxV)
		count += dfs(node.Right, maxV)
		return count
	}
	return dfs(root, root.Val)
}

func main() {
	// root = [3,1,4,3,null,1,5]
	node6 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 1, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 4, Left: node5, Right: node6}
	node2 := &TreeNode{Val: 1, Left: node4, Right: nil}
	head := &TreeNode{Val: 3, Left: node2, Right: node3}
	result := goodNodes(head)

	fmt.Println("Nums goodNodes:", result)
}
