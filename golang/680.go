//go:build ignore

package main

// Given a string s, return true if the s can be palindrome after deleting at most one character from it.

// Example 1:
//
// Input: s = "aba"
// Output: true
// Example 2:
//
// Input: s = "abca"
// Output: true
// Explanation: You could delete the character 'c'.
// Example 3:
//
// Input: s = "abc"
// Output: false
//
//
// Constraints:
//
// 1 <= s.length <= 105
// s consists of lowercase English letters

func validPalindrome(s string) bool {
	isPalindrome := func(str string) bool {
		l, r := 0, len(str)-1

		for l < r {
			if str[l] != str[r] {
				return false
			}
			l++
			r--
		}

		return true
	}
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			skipL := s[l+1 : r+1]
			slipR := s[l:r]

			return isPalindrome(skipL) || isPalindrome(slipR)

		}

		l++
		r--
	}

	return true
}

func main() {
}
