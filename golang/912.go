//go:build ignore

package main

// Given an array of integers nums, sort the array in ascending order and return it.
//
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.
//
//
//
// Example 1:
//
// Input: nums = [5,2,3,1]
// Output: [1,2,3,5]
// Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).
// Example 2:
//
// Input: nums = [5,1,1,2,0,0]
// Output: [0,0,1,1,2,5]
// Explanation: Note that the values of nums are not necessarily unique.
//
//
// Constraints:
//
// 1 <= nums.length <= 5 * 104
// -5 * 104 <= nums[i] <= 5 * 104

func sortArray(nums []int) []int {
	// merge sort

	var merge func(arr []int, l, m, r int)
	merge = func(arr []int, l, m, r int) {
		left := make([]int, m-l+1)
		right := make([]int, r-m)
		copy(left, arr[l:m+1])
		copy(right, arr[m+1:r+1])

		i, j, k := l, 0, 0
		for j < len(left) && k < len(right) {
			if left[j] >= right[k] {
				arr[i] = right[k]
				k++
			} else {
				arr[i] = left[j]
				j++
			}
			i++
		}
		for j < len(left) {
			arr[i] = left[j]
			i++
			j++
		}
		for k < len(right) {
			arr[i] = right[k]
			i++
			k++
		}
	}

	var mergeSort func(arr []int, l, r int)
	mergeSort = func(arr []int, l, r int) {
		if l >= r {
			return
		}

		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
}
