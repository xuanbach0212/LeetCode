//go:build ignore

package main

import (
	"fmt"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
//
//
// Example 1:
//
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:
//
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:
//
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
//

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		cl := s[left]
		cr := s[right]

		if !isAlphaNum(cl) {
			left++
			continue
		}

		if !isAlphaNum(cr) {
			right--
			continue
		}

		if toLower(cl) != toLower(cr) {
			return false
		}

		left++
		right--

	}
	return true
}

func isAlphaNum(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}

func main() {
	fmt.Println("isPalindrome is:", isPalindrome("bach"))
}
