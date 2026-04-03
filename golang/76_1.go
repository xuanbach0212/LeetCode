//go:build ignore

package main

import (
	"fmt"
	"math"
)

// Given two strings s and t of lengths m and n respectively, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.
//
//
//
// Example 1:
//
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
// Example 2:
//
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.
// Example 3:
//
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

func minWindow(s, t string) string {
	if t == "" {
		return ""
	}
	res := ""
	nRes := math.MaxInt
	window := map[byte]int{}
	countT := map[byte]int{}

	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}
	match, need := 0, len(countT)

	l := 0
	for r := 0; r < len(s); r++ {
		window[s[r]]++

		if window[s[r]] == countT[s[r]] {
			match++
		}

		for match == need {
			sub := s[l : r+1]
			if len(sub) < nRes {
				nRes = len(sub)
				res = sub
			}

			window[s[l]]--
			if countT[s[l]] > 0 && countT[s[l]] > window[s[l]] {
				match--
			}
			l++
		}
	}

	return res
}

func main() {
	fmt.Println("minWindow is: ", minWindow("ADOBECODEBANC", "ABC"))
}
