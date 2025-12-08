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

func getSuccessor(curr *TreeNode) *TreeNode {
	curr = curr.Right
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		succ := getSuccessor(root)
		root.Val = succ.Val
		root.Right = deleteNode(root.Right, succ.Val)
	}

	return root
}

func main() {
	// root = [5,3,6,2,4,null,7], key = 3
	vals := []*int{
		intPtr(5), intPtr(3), intPtr(6), intPtr(2), intPtr(4), nil, intPtr(7),
	}
	root := BuildTree(vals)
	result := deleteNode(root, 3)

	fmt.Println("deleteNode", result.Val)
}
