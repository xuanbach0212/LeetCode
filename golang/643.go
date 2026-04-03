//go:build ignore

package main

import (
	"fmt"
	"math"
)

// You are given an integer array nums consisting of n elements, and an integer k.
//
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
//
//
//
// Example 1:
//
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
// Example 2:
//
// Input: nums = [5], k = 1
// Output: 5.00000
//
//
// Constraints:
//
// n == nums.length
// 1 <= k <= n <= 105
// -104 <= nums[i] <= 104

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxAVG := -math.MaxFloat64
	for r, num := range nums {
		sum += num
		if r < k-1 {
			continue
		}

		maxAVG = max(float64(sum)/float64(k), maxAVG)
		sum -= nums[r-k+1]
	}
	return maxAVG
}

func main() {
	fmt.Println("findMaxAverage is: ", findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
