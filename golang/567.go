//go:build ignore

package main

import (
	"fmt"
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
	if len(s1) > len(s2) {
		return false
	}

	s1Arr := make([]int, 26)
	s2Arr := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		s1Arr[s1[i]-'a']++
		s2Arr[s2[i]-'a']++
	}

	match := 0
	for i := 0; i < 26; i++ {
		if s1Arr[i] == s2Arr[i] {
			match += 1
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if match == 26 {
			return true
		}

		idx := s2[r] - 'a'
		s2Arr[idx]++
		if s1Arr[idx] == s2Arr[idx] {
			match++
		} else if s1Arr[idx]+1 == s2Arr[idx] {
			match--
		}

		idx = s2[l] - 'a'
		s2Arr[idx]--
		if s1Arr[idx] == s2Arr[idx] {
			match++
		} else if s1Arr[idx]-1 == s2Arr[idx] {
			match--
		}
		l++
	}

	return match == 26
}

func main() {
	fmt.Println("checkInclusion is:", checkInclusion("ab", "eidboaoo"))
}
