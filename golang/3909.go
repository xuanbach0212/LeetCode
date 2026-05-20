//go:build ignore

package main

import "fmt"

func compareBitonicSums(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	peakIndex := left + (right-left)/2

	ascendingSum, descendingSum := 0, 0
	for i := 0; i <= peakIndex; i++ {
		ascendingSum += nums[i]
	}

	for i := peakIndex; i < len(nums); i++ {
		descendingSum += nums[i]
	}

	if ascendingSum > descendingSum {
		return 0
	} else if descendingSum > ascendingSum {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(compareBitonicSums([]int{1, 3, 2, 1}))
	fmt.Println(compareBitonicSums([]int{2, 4, 5, 2}))
	fmt.Println(compareBitonicSums([]int{1, 2, 4, 3}))
}
