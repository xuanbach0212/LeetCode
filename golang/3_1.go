//go:build ignore

package main

import (
	"fmt"
)

// Given a string s, find the length of the longest substring without duplicate characters.
//
//
//
// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	l := 0
	res := 0
	count := make(map[byte]int, len(s))
	for r := 0; r < len(s); r++ {
		count[s[r]]++

		for count[s[r]] > 1 {
			count[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}

func main() {
	fmt.Println("lengthOfLongestSubstring is:", lengthOfLongestSubstring("pwwkew"))
}
