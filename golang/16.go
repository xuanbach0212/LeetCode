//go:build ignore

package main

import "sort"

// Given an integer array nums of length n and an integer target, find three integers at distinct indices in nums such that the sum is closest to target.
//
// Return the sum of the three integers.
//
// You may assume that each input would have exactly one solution.
//
//
//
// Example 1:
//
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// Example 2:
//
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
//
//
// Constraints:
//
// 3 <= nums.length <= 500
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104

func absInt(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closedSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			currSum := nums[i] + nums[l] + nums[r]
			if absInt(target-currSum) < absInt(target-closedSum) {
				closedSum = currSum
			}

			if currSum < target {
				l++
			} else if currSum > target {
				r--
			} else {
				return target
			}
		}
	}

	return closedSum
}

func main() {
}
