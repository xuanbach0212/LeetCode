//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
//
//
// Example 1:
//
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// Example 2:
//
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
// Example 3:
//
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.
//
// [-4,-1,-1,0,1,2]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := range nums {
		// handle if same value of first idx so keeep going
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// looop with 2 pointer, first pointer is i
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]

			if sum > 0 {
				r--
				continue
			}

			if sum < 0 {
				l++
				continue
			}

			result = append(result, []int{nums[l], nums[i], nums[r]})
			l++

			// handle not using the same value l for distinct
			for nums[l] == nums[l-1] && l < r {
				l++
			}
		}
	}

	return result
}

func main() {
	fmt.Println("three Sum is:", threeSum([]int{0, 0, 0, 0}))
}
