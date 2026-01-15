//go:build ignore

package main

import (
	"fmt"
)

// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.
//
// Return the max sliding window.
//
//
//
// Example 1:
//
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
// Example 2:
//
// Input: nums = [1], k = 1
// Output: [1]

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	queue := []int{}
	l, r := 0, 0

	for r < len(nums) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[r] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, r)

		if queue[0] < l {
			queue = queue[1:]
		}

		if r-l+1 == k {
			result = append(result, nums[queue[0]])
			l++
		}
		r++
	}
	return result
}

func main() {
	fmt.Println("maxSlidingWindow is: ", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
