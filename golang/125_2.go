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

func main() {
	fmt.Println("isPalindrome is:", isPalindrome("bach"))
}
