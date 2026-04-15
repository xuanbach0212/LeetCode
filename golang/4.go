//go:build ignore

package main

import "math"

// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
//
//
//
// Example 1:
//
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:
//
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
//
// Constraints:
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(A) + len(B)
	half := total / 2

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)
	for l <= r {
		p1 := (l + r) / 2
		p2 := half - p1

		l1 := math.MinInt
		if p1 > 0 {
			l1 = A[p1-1]
		}
		r1 := math.MaxInt
		if p1 < len(A) {
			r1 = A[p1]
		}

		l2 := math.MinInt
		if p2 > 0 {
			l2 = B[p2-1]
		}
		r2 := math.MaxInt
		if p2 < len(B) {
			r2 = B[p2]
		}

		if l1 <= r2 && l2 <= r1 {
			if (total)%2 == 1 {
				return float64(min(r1, r2))
			}
			return float64(max(l1, l2)+min(r1, r2)) / 2.0
		} else if l1 > r2 {
			r = p1 - 1
		} else {
			l = p1 + 1
		}

	}

	return 0
}

func main() {
}
