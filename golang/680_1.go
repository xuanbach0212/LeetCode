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

func isValid(c byte) bool {
	if (c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if c >= 65 && c <= 90 {
		return c + 32
	}
	return c
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !isValid(s[l]) {
			l++
			continue
		}

		if !isValid(s[r]) {
			r--
			continue
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			skipL := s[l+1 : r+1]
			skipR := s[l:r]

			return isPalindrome(skipL) || isPalindrome(skipR)
		}
		l++
		r--
	}
	return true
}

func main() {
}
