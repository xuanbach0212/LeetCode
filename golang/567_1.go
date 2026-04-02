//go:build ignore

package main

import (
	"fmt"
	"slices"
)

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of s2.
//
//
//
// Example 1:
//
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:
//
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false

func checkInclusion(s1 string, s2 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	if n1 > n2 {
		return false
	}

	count1 := make([]int, 26)
	count2 := make([]int, 26)

	for i := 0; i < n1; i++ {
		count1[s1[i]-'a'] += 1
		count2[s2[i]-'a'] += 1
	}

	if slices.Equal(count1, count2) {
		return true
	}

	for j := n1; j < n2; j++ {
		count2[s2[j]-'a'] += 1
		count2[s2[j-n1]-'a'] -= 1

		if slices.Equal(count1, count2) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("checkInclusion is:", checkInclusion("ab", "eidboaoo"))
}
