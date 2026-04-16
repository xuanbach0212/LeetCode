//go:build ignore

package main

import (
	"fmt"
)

// Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.
//
// Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2. Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.
//
//
//
// Example 1:
//
// Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
// Output: [2,2,2,1,4,3,3,9,6,7,19]
// Example 2:
//
// Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
// Output: [22,28,8,6,17,44]
//
//
// Constraints:
//
// 1 <= arr1.length, arr2.length <= 1000
// 0 <= arr1[i], arr2[i] <= 1000
// All the elements of arr2 are distinct.
// Each arr2[i] is in arr1.

func relativeSortArray(arr1 []int, arr2 []int) []int {
	maxV := 0
	for _, num := range arr1 {
		if num > maxV {
			maxV = num
		}
	}

	count := make([]int, maxV+1)
	for _, num := range arr1 {
		count[num]++
	}

	var res []int
	for _, num := range arr2 {
		for count[num] > 0 {
			res = append(res, num)
			count[num]--
		}
	}

	for num := 0; num < maxV+1; num++ {
		for count[num] > 0 {
			res = append(res, num)
			count[num]--
		}
	}

	return res
}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
