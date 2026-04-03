//go:build ignore

package main

import (
	"fmt"
)

// Given an integer array nums and an integer k, return the number of good subarrays of nums.
//
// A good array is an array where the number of different integers in that array is exactly k.
//
// For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
// A subarray is a contiguous part of an array.
//
//
//
// Example 1:
//
// Input: nums = [1,2,1,2,3], k = 2
// Output: 7
// Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
// Example 2:
//
// Input: nums = [1,2,1,3,4], k = 3
// Output: 3
// Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
//
//
// Constraints:
//
// 1 <= nums.length <= 2 * 104
// 1 <= nums[i], k <= nums.length

func subarraysWithKDistinct(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {
	res := 0
	distict := 0
	window := make([]int, len(nums)+1)
	l := 0
	for r, num := range nums {
		if window[num] == 0 {
			distict++
		}
		window[num]++

		for distict > k {
			window[nums[l]]--
			if window[nums[l]] == 0 {
				distict--
			}
			l++
		}

		res += r - l + 1
	}
	return res
}

func main() {
	fmt.Println("subarraysWithKDistinct is: ", subarraysWithKDistinct([]int{2, 1, 2, 1, 2}, 2))
}
