//go:build ignore

package main

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
//
// Example 1:
//
// Input: nums = [3,2,3]
// Output: [3]
// Example 2:
//
// Input: nums = [1]
// Output: [1]
// Example 3:
//
// Input: nums = [1,2]
// Output: [1,2]
//
// Constraints:
//
// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109
//
// Follow up: Could you solve the problem in linear time and in O(1) space?

func count(nums []int, el int) int {
	count := 0
	for _, num := range nums {
		if num == el {
			count += 1
		}
	}
	return count
}

func majorityElement(nums []int) []int {
	m := map[int]int{}
	res := []int{}

	for _, num := range nums {
		m[num]++

		if len(m) > 2 {
			for k := range m {
				m[k] -= 1
				if m[k] == 0 {
					delete(m, k)
				}
			}
		}
	}

	for k := range m {
		if count(nums, k) > len(nums)/3 {
			res = append(res, k)
		}
	}

	return res
}

func main() {
}
