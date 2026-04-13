//go:build ignore

package main

import "fmt"

// You are given an m x n integer matrix matrix with the following two properties:
//
// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.
//
// You must write a solution in O(log(m * n)) time complexity.
//
//
//
// Example 1:
//
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true
// Example 2:
//
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false
//
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -104 <= matrix[i][j], target <= 104

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1

	for l <= r {
		mid := l + (r-l)/2

		if matrix[mid][0] > target {
			r = mid - 1
		} else if matrix[mid][len(matrix[0])-1] < target {
			l = mid + 1
		} else {
			break
		}
	}

	if !(l <= r) {
		return false
	}

	lt, rt := 0, len(matrix[l])-1
	row := (l + r) / 2
	for lt <= rt {
		mt := lt + (rt-lt)/2

		if matrix[row][mt] > target {
			rt = mt - 1
		} else if matrix[row][mt] < target {
			lt = mt + 1
		} else {
			return true
		}

	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 6, 20}, {23, 30, 34, 60}}, 3))
}
