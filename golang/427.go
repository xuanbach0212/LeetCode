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

// Given a n * n matrix grid of 0's and 1's only. We want to represent grid with a Quad-Tree.
//
// Return the root of the Quad-Tree representing grid.
//
// A Quad-Tree is a tree data structure in which each internal node has exactly four children. Besides, each node has two attributes:
//
// val: True if the node represents a grid of 1's or False if the node represents a grid of 0's. Notice that you can assign the val to True or False when isLeaf is False, and both are accepted in the answer.
// isLeaf: True if the node is a leaf node on the tree or False if the node has four children.
// class Node {
//     public boolean val;
//     public boolean isLeaf;
//     public Node topLeft;
//     public Node topRight;
//     public Node bottomLeft;
//     public Node bottomRight;
// }
// We can construct a Quad-Tree from a two-dimensional area using the following steps:
//
// If the current grid has the same value (i.e all 1's or all 0's) set isLeaf True and set val to the value of the grid and set the four children to Null and stop.
// If the current grid has different values, set isLeaf to False and set val to any value and divide the current grid into four sub-grids as shown in the photo.
// Recurse for each of the children with the proper sub-grid.
//
// If you want to know more about the Quad-Tree, you can refer to the wiki.
//
// Quad-Tree format:
//
// You don't need to read this section for solving the problem. This is only if you want to understand the output format here. The output represents the serialized format of a Quad-Tree using level order traversal, where null signifies a path terminator where no node exists below.
//
// It is very similar to the serialization of the binary tree. The only difference is that the node is represented as a list [isLeaf, val].
//
// If the value of isLeaf or val is True we represent it as 1 in the list [isLeaf, val] and if the value of isLeaf or val is False we represent it as 0.
//
//
//
// Example 1:
//
//
// Input: grid = [[0,1],[1,0]]
// Output: [[0,1],[1,0],[1,1],[1,1],[1,0]]
// Explanation: The explanation of this example is shown below:
// Notice that 0 represents False and 1 represents True in the photo representing the Quad-Tree.
//
// Example 2:
//
//
//
// Input: grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
// Output: [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
// Explanation: All values in the grid are not the same. We divide the grid into four sub-grids.
// The topLeft, bottomLeft and bottomRight each has the same value.
// The topRight have different values so we divide it into 4 sub-grids where each has the same value.
// Explanation is shown in the photo below:
//
//
//
// Constraints:
//
// n == grid.length == grid[i].length
// n == 2x where 0 <= x <= 6

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func isLeafGrid(subGrid [][]int) bool {
	toCompare := subGrid[0][0]
	for i := range subGrid {
		for j := range subGrid {
			if subGrid[i][j] != toCompare {
				return false
			}
		}
	}
	return true
}

func construct(grid [][]int) *Node {
	// found len grid
	// each sub-grid with divide
	// check each sub grid to detect is leaf or not
	// if not do all the same again

	lenGrid := len(grid)
	lenOfEachSub := lenGrid / 2
	res := &Node{
		IsLeaf: true,
		Val:    grid[0][0] == 1,
	}

	if isLeafGrid(grid) {
		return res
	}

	topLeftGrid := [][]int{}
	topRightGrid := [][]int{}
	bottomLeftGrid := [][]int{}
	bottomRightGrid := [][]int{}

	for idx, row := range grid {
		// idx from 0 to lenOfEachSub
		if idx < lenOfEachSub {
			topLeftGrid = append(topLeftGrid, row[:lenOfEachSub])
			topRightGrid = append(topRightGrid, row[lenOfEachSub:])
		} else {
			bottomLeftGrid = append(bottomLeftGrid, row[:lenOfEachSub])
			bottomRightGrid = append(bottomRightGrid, row[lenOfEachSub:])
		}
	}

	res.IsLeaf = false
	res.Val = true
	res.TopLeft = construct(topLeftGrid)
	res.TopRight = construct(topRightGrid)
	res.BottomLeft = construct(bottomLeftGrid)
	res.BottomRight = construct(bottomRightGrid)

	return res
}

// follow up
func contructFollowUp(grid [][]int) *Node {
	return helper(grid, 0, 0, len(grid))
}

// use r, c, s to decrease space to decrease space complexity
func helper(grid [][]int, r, c, s int) *Node {
	val := grid[r][c]
	isLeaf := true

	res := &Node{
		IsLeaf: true,
		Val:    val == 1,
	}

	// check if leaf
	for i := r; i < r+s; i++ {
		for j := c; j < c+s; j++ {
			if grid[i][j] != val {
				isLeaf = false
				break
			}
		}
		if !isLeaf {
			break
		}
	}

	if isLeaf {
		return res
	}

	half := s / 2
	res.IsLeaf = false
	res.Val = false
	res.TopLeft = helper(grid, r, c, half)
	res.TopRight = helper(grid, r, c+half, half)
	res.BottomLeft = helper(grid, r+half, c, half)
	res.BottomRight = helper(grid, r+half, c+half, half)

	return res
}

func main() {
	// root = [1,2,3,null,5,null,4]
	//	vals := []*int{
	//		intPtr(1), intPtr(2), intPtr(3), nil, intPtr(5), nil, intPtr(4),
	//	}
	//	root := BuildTree(vals)
	//	result := construct(root)
	//
	//	for _, v := range result {
	//		fmt.Println(" rightSideView is:", v)
	//	}
}
