//go:build ignore

package main

import (
	"fmt"
)

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
//
//
//
// Example 1:
//
//
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
// Example 2:
//
// Input: height = [4,2,0,3,2,5]
// Output: 9

func trap(height []int) int {
	l, r := 0, len(height)-1
	maxL, maxR := 0, 0
	res := 0
	for l < r {
		if height[l] < height[r] {
			maxL = max(maxL, height[l])
			res += maxL - height[l]
			l++
		} else {
			maxR = max(maxR, height[r])
			res += maxR - height[r]
			r--
		}
	}

	return res
}

func main() {
	fmt.Println("trap is:", trap([]int{4, 2, 0, 3, 2, 5}))
}
