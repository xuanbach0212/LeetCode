//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val >= maxVal {
			res = 1
		}
		maxVal = max(node.Val, maxVal)
		res += dfs(node.Left, maxVal)
		res += dfs(node.Right, maxVal)
		return res
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
