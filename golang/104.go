//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {
	// root = [3,9,20,null,null,15,7]
	node5 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 20, Left: node4, Right: node5}
	node2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	head := &TreeNode{Val: 3, Left: node2, Right: node3}
	result := maxDepth(head)

	fmt.Println("Maximum Depth:", result)
}
