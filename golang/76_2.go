//go:build ignore

package main

import (
	"fmt"
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
	if len(s) < len(t) {
		return ""
	}
	windowmap := [128]int{}
	tmap := [128]int{}

	need := 0
	for i := 0; i < len(t); i++ {
		if tmap[t[i]] == 0 {
			need++
		}
		tmap[t[i]] += 1
	}

	have := 0
	l := 0
	start := l
	minlen := len(s) + 1
	for r := 0; r < len(s); r++ {
		c := s[r]
		windowmap[c]++

		if tmap[c] > 0 && tmap[c] == windowmap[c] {
			have++
		}

		for have == need {
			if r-l+1 < minlen {
				minlen = r - l + 1
				start = l
			}

			windowmap[s[l]]--
			if tmap[s[l]] > 0 && tmap[s[l]] > windowmap[s[l]] {
				have--
			}
			l++
		}

	}

	if minlen > len(s) {
		return ""
	}

	return s[start : start+minlen]
}

func main() {
	fmt.Println("minWindow is: ", minWindow("bdab", "bd"))
}
