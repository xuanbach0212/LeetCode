//go:build ignore

package main

// Given a 2D matrix matrix, handle multiple queries of the following type:
//
// Calculate the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
// Implement the NumMatrix class:
//
// NumMatrix(int[][] matrix) Initializes the object with the integer matrix matrix.
// int sumRegion(int row1, int col1, int row2, int col2) Returns the sum of the elements of matrix inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).
// You must design an algorithm where sumRegion works on O(1) time complexity.
//
// [0, 0,  0,  0]
// [0, 3,  3,  4]
// [0, 8, 14, 18]
// [0, 9, 17, 21]
//
//
// Input
// ["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
// [[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
// Output
// [null, 8, 11, 12]
//
// Explanation
// NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
// numMatrix.sumRegion(2, 1, 4, 3); // return 8 (i.e sum of the red rectangle)
// numMatrix.sumRegion(1, 1, 2, 2); // return 11 (i.e sum of the green rectangle)
// numMatrix.sumRegion(1, 2, 2, 4); // return 12 (i.e sum of the blue rectangle)
//
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// -104 <= matrix[i][j] <= 104
// 0 <= row1 <= row2 < m
// 0 <= col1 <= col2 < n
// At most 104 calls will be made to sumRegion.

type NumMatrix struct {
	prefixMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix) + 1
	n := len(matrix[0]) + 1
	prefixMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		prefixMatrix[i] = make([]int, n)
	}
	for i, r := range matrix {
		prefixSum := 0
		for j, c := range r {
			prefixSum += c
			prefixMatrix[i+1][j+1] = prefixSum + prefixMatrix[i][j+1]
		}
	}
	return NumMatrix{prefixMatrix: prefixMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	col1++
	row2++
	col2++
	topLeft := this.prefixMatrix[row1-1][col1-1]
	above := this.prefixMatrix[row1-1][col2]
	bottomLeft := this.prefixMatrix[row2][col1-1]
	bottomRight := this.prefixMatrix[row2][col2]
	return bottomRight - above - bottomLeft + topLeft
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
}
