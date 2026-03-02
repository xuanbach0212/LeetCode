//go:build ignore

package main

import (
	"fmt"
)

// There is an integer array nums sorted in ascending order (with distinct values).
//
// Prior to being passed to your function, nums is possibly left rotated at an unknown index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be left rotated by 3 indices and become [4,5,6,7,0,1,2].
//
// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
//
// Example 1:
//
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:
//
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
// Example 3:
//
// Input: nums = [1], target = 0
// Output: -1

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (r + l) / 2
		if m == target {
			return m
		}

		if nums[l] <= nums[m] {
			if nums[l] > target {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if nums[l+1] < target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println("search is: ", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
