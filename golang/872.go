//go:build ignore

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(root *TreeNode, leaf *[]int)
	dfs = func(root *TreeNode, leaf *[]int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			*leaf = append(*leaf, root.Val)
			return
		}
		dfs(root.Left, leaf)
		dfs(root.Right, leaf)
	}
	leaf1, leaf2 := []int{}, []int{}
	dfs(root1, &leaf1)
	dfs(root2, &leaf2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i, v := range leaf1 {
		if v != leaf2[i] {
			return false
		}
	}
	return true
}

func main() {
	// // oot1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
	// node5 := &TreeNode{Val: 7, Left: nil, Right: nil}
	// node4 := &TreeNode{Val: 15, Left: nil, Right: nil}
	// node3 := &TreeNode{Val: 20, Left: node4, Right: node5}
	// node2 := &TreeNode{Val: 9, Left: nil, Right: nil}
	// head := &TreeNode{Val: 3, Left: node2, Right: node3}
	// result := maxDepth(head)
	//
	// fmt.Println("Maximum Depth:", result)
}
