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

func deleteNodeIterative(root *TreeNode, key int) *TreeNode {
	// i think we need to find parent then find the suitable for replace. father -> suitable

	var father *TreeNode = nil
	curr := root

	// find father and find the node need to be delete
	for curr != nil && curr.Val != key {
		father = curr
		if key > curr.Val {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

	// didnt found key
	if curr == nil {
		return root
	}

	// find suitable node to replace
	// we have 3 case, if child has one or none, if child has two
	// if first then replace
	// if two then in-order successor

	if curr.Left == nil || curr.Right == nil {
		var newCurr *TreeNode
		if curr.Left != nil {
			newCurr = curr.Left
		} else {
			newCurr = curr.Right
		}

		if father == nil {
			// this one mean the root is the one to delete
			return newCurr
		}

		if father.Left == curr {
			father.Left = newCurr
		} else {
			father.Right = newCurr
		}

	} else {
		// in-order successor
		// find then connect successor child and father
		successor := curr.Right
		fatherSuccessor := curr

		for successor.Left != nil {
			fatherSuccessor = successor
			successor = successor.Left
		}

		curr.Val = successor.Val

		// this one tricky because, if successor is curr dont have child in the left so. parrent of successor must connect to successor's right child
		if fatherSuccessor == curr {
			fatherSuccessor.Right = successor.Right
		} else {
			fatherSuccessor.Left = successor.Right
		}

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
