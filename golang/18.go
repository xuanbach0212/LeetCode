//go:build ignore

package main

import (
	"sort"
)

// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
//
// 0 <= a, b, c, d < n
// a, b, c, and d are distinct.
// nums[a] + nums[b] + nums[c] + nums[d] == target
// You may return the answer in any order.
//
//
//
// Example 1:
//
// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// Example 2:
//
// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]
//
//
// Constraints:
//
// 1 <= nums.length <= 200
// -109 <= nums[i] <= 109
// -109 <= target <= 109

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	quad := []int{}

	var helper func(k int, target int, start int)
	helper = func(k, target, start int) {
		if k == 2 {
			l, r := start, len(nums)-1
			for l < r {
				if nums[l]+nums[r] > target {
					r--
				} else if nums[l]+nums[r] < target {
					l++
				} else {
					tmp := make([]int, len(quad))
					copy(tmp, quad)
					tmp = append(tmp, nums[l], nums[r])
					res = append(res, tmp)

					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
			return
		}
		for i := start; i < len(nums)-k+1; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			quad = append(quad, nums[i])
			helper(k-1, target-nums[i], i+1)
			quad = quad[:len(quad)-1]
		}
	}
	helper(4, target, 0)
	return res
}

func main() {
}
