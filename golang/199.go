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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	rightSideNodes := []int{}

	for len(queue) > 0 {
		var rightSide *TreeNode
		levelSize := len(queue)
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				rightSide = node
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		if rightSide != nil {
			rightSideNodes = append(rightSideNodes, rightSide.Val)
		}
	}

	return rightSideNodes
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
