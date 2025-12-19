//go:build ignore

package main

import (
	"fmt"
)

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
//
// Return the length of the longest substring containing the same letter you can get after performing the above operations.
//
//
//
// Example 1:
//
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
// Example 2:
//
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.
// AABABBBA

func characterReplacement(s string, k int) int {
	res := 0
	charSet := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		charSet[s[i]] = true
	}

	for c := range charSet {
		count, l := 0, 0
		for r := 0; r < len(s); r++ {
			if s[r] == c {
				count++
			}

			for (r-l+1)-count > k {
				if s[l] == c {
					count--
				}
				l++
			}
			res = max(res, r-l+1)
		}
	}

	return res
}

func main() {
	fmt.Println("characterReplacement is:", characterReplacement("ABBB", 2))
}
