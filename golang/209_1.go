//go:build ignore

package main

import "math"

// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.
//
//
//
// Example 1:
//
// Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2
// Explanation: The subarray [4,3] has the minimal length under the problem constraint.
// Example 2:
//
// Input: target = 4, nums = [1,4,4]
// Output: 1
// Example 3:
//
// Input: target = 11, nums = [1,1,1,1,1,1,1,1]
// Output: 0
//
//
// Constraints:
//
// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 104
//
//
// Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	res := math.MaxInt
	sum := 0

	for r < len(nums) {
		sum += nums[r]
		for sum >= target && l <= r {
			res = min(r-l+1, res)
			sum -= nums[l]
			l++
		}
		r++
	}

	if res == math.MaxInt {
		return 0
	}
	return res
}

func main() {
}
