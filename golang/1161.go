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

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxLevel := 1
	maxSum := root.Val
	level := 1
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if maxSum < sum {
			maxSum = sum
			maxLevel = level
		}
		level += 1
	}

	return maxLevel
}

func main() {
	// root = [1,7,0,7,-8,null,null]
	vals := []*int{
		intPtr(1), intPtr(7), intPtr(0), intPtr(7), intPtr(-8), nil, nil,
	}
	root := BuildTree(vals)
	result := maxLevelSum(root)

	fmt.Println(" rightSideView is:", result)
}
