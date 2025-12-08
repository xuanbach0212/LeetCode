//go:build ignore

package main

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

func searchBST(root *TreeNode, val int) *TreeNode {
	for root == nil {
		return nil
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	return root
}

func main() {
	// root = [4,2,7,1,3]
	vals := []*int{
		intPtr(4), intPtr(2), intPtr(7), intPtr(1), intPtr(3),
	}
	root := BuildTree(vals)
	result := searchBST(root, 2)
}
