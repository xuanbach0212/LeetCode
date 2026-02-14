//go:build ignore

package main

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
//
//
// Example 1:
//
// Input: s = "anagram", t = "nagaram"
//
// Output: true
//
// Example 2:
//
// Input: s = "rat", t = "car"
//
// Output: false
//
//
//
// Constraints:
//
// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.
//
//
// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, el := range count {
		if el != 0 {
			return false
		}
	}

	return true
}

func main() {
}
