//go:build ignore

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	maxLength := 0
	var dfs func(node *TreeNode, goLeft bool, length int)
	dfs = func(node *TreeNode, goLeft bool, length int) {
		if node == nil {
			return
		}

		maxLength = max(length, maxLength)

		if goLeft {
			dfs(node.Left, false, length+1)
			dfs(node.Right, true, 1)
		} else {
			dfs(node.Right, true, length+1)
			dfs(node.Left, false, 1)
		}
	}

	dfs(root.Left, false, 1)
	dfs(root.Right, true, 1)

	return maxLength
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

func main() {
	vals := []*int{
		intPtr(1), nil, intPtr(1), intPtr(1), intPtr(1), nil, nil, intPtr(1), intPtr(1), nil, intPtr(1), nil, nil, nil, intPtr(1),
	}
	root := BuildTree(vals)
	result := longestZigZag(root)

	fmt.Println("logest zigzag is:", result)
}
