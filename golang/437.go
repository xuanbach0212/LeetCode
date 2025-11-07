//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(node *TreeNode, cnt map[int]int, prefixSum int) int
	cnt := map[int]int{
		0: 1,
	}
	dfs = func(node *TreeNode, cnt map[int]int, prefixSum int) int {
		if node == nil {
			return 0
		}
		prefixSum += node.Val
		res := cnt[prefixSum-targetSum]
		cnt[prefixSum]++
		res += dfs(node.Left, cnt, prefixSum)
		res += dfs(node.Right, cnt, prefixSum)
		cnt[prefixSum]--
		return res
	}
	return dfs(root, cnt, 0)
}

func main() {
	// root = [10,5,-3,3,2,null,11,3,-2,null,1]
	// targetSum = 8
	node11 := &TreeNode{Val: 1}
	node10 := &TreeNode{Val: -2}
	node9 := &TreeNode{Val: 3}
	node8 := &TreeNode{Val: 11}
	node7 := &TreeNode{Val: 2, Right: node11}
	node6 := &TreeNode{Val: 3, Left: node9, Right: node10}
	node5 := &TreeNode{Val: -3, Right: node8}
	node4 := &TreeNode{Val: 5, Left: node6, Right: node7}
	head := &TreeNode{Val: 10, Left: node4, Right: node5}

	result := pathSum(head, 8)

	fmt.Println("path Sum is:", result)
}
