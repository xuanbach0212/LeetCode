//go:build ignore

package main

import "fmt"

// Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.
//
// A subarray is a contiguous part of an array.
//
//
//
// Example 1:
//
// Input: nums = [4,5,0,-2,-3,1], k = 5
// Output: 7
// Explanation: There are 7 subarrays with a sum divisible by k = 5:
// [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
// Example 2:
//
// Input: nums = [5], k = 9
// Output: 0
//
//
// Constraints:
//
// 1 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
// 2 <= k <= 104

func subarraysDivByK(nums []int, k int) int {
	prefixSum, res := 0, 0
	prefixCnt := make(map[int]int)
	prefixCnt[0] = 1

	for _, n := range nums {
		prefixSum += n
		remain := prefixSum % k
		if remain < 0 {
			remain += k
		}
		res += prefixCnt[remain]
		prefixCnt[remain]++
	}

	return res
}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	fmt.Println(subarraysDivByK([]int{5}, 9))
}
