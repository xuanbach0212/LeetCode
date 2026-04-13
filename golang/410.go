//go:build ignore

package main

// Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.
//
// Return the minimized largest sum of the split.
//
// A subarray is a contiguous part of the array.
//
//
//
// Example 1:
//
// Input: nums = [7,2,5,10,8], k = 2
// Output: 18
// Explanation: There are four ways to split nums into two subarrays.
// The best way is to split it into [7,2,5] and [10,8], where the largest sum among the two subarrays is only 18.
// Example 2:
//
// Input: nums = [1,2,3,4,5], k = 2
// Output: 9
// Explanation: There are four ways to split nums into two subarrays.
// The best way is to split it into [1,2,3] and [4,5], where the largest sum among the two subarrays is only 9.
//
//
// Constraints:
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 106
// 1 <= k <= min(50, nums.length)

func canSplit(nums []int, mid int, k int) bool {
	count := 1
	total := 0

	for _, num := range nums {
		if total+num > mid {
			count++
			total = num
			if count > k {
				return false
			}
		} else {
			total += num
		}
	}
	return true
}

func splitArray(nums []int, k int) int {
	l, r := 0, 0

	for _, num := range nums {
		r += num
		if l < num {
			l = num
		}
	}

	res := 0
	for l <= r {
		m := l + (r-l)/2

		if canSplit(nums, m, k) {
			res = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return res
}

// total: 32
// 7  25
// 9  23
// 14 18
// 24 8

func main() {
}
