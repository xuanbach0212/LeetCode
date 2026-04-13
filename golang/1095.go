//go:build ignore

package main

// (This problem is an interactive problem.)
//
// You may recall that an array arr is a mountain array if and only if:
//
// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
// Given a mountain array mountainArr, return the minimum index such that mountainArr.get(index) == target. If such an index does not exist, return -1.
//
// You cannot access the mountain array directly. You may only access the array using a MountainArray interface:
//
// MountainArray.get(k) returns the element of the array at index k (0-indexed).
// MountainArray.length() returns the length of the array.
// Submissions making more than 100 calls to MountainArray.get will be judged Wrong Answer. Also, any solutions that attempt to circumvent the judge will result in disqualification.
//
//
//
// Example 1:
//
// Input: mountainArr = [1,2,3,4,5,3,1], target = 3
// Output: 2
// Explanation: 3 exists in the array, at index=2 and index=5. Return the minimum index, which is 2.
// Example 2:
//
// Input: mountainArr = [0,1,2,4,2,1], target = 3
// Output: -1
// Explanation: 3 does not exist in the array, so we return -1.
//
//
// Constraints:
//
// 3 <= mountainArr.length() <= 104
// 0 <= target <= 109
// 0 <= mountainArr.get(index) <= 109

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	N := mountainArr.length()

	l, r := 0, N-1
	// find the peak
	peak := 1
	for l <= r {
		m := (l + r) / 2
		left := mountainArr.get(m - 1)
		right := mountainArr.get(m + 1)
		mid := mountainArr.get(m)

		if left < mid && mid < right {
			l = m + 1
		} else if left > mid && mid > right {
			r = m - 1
		} else {
			peak = m
			break
		}
	}

	// find target in the left side
	for l, r = 0, peak-1; l <= r; {
		m := (l + r) / 2
		val := mountainArr.get(m)

		if val > target {
			r = m - 1
		} else if val < target {
			l = m + 1
		} else {
			return m
		}
	}

	// find target in the right side
	for l, r = peak, N-1; l <= r; {
		m := (l + r) / 2
		val := mountainArr.get(m)

		if val > target {
			l = m + 1
		} else if val < target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}

func main() {
}
