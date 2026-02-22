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
	l := 0
	r := len(s) - 1
	for l < r {
		if !isValid(s[l]) {
			l++
			continue
		}
		if !isValid(s[r]) {
			r--
			continue
		}
		if converUppercaseToLowercase(s[l]) != converUppercaseToLowercase(s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isValid(char byte) bool {
	if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) || (char >= 48 && char <= 57) {
		return true
	}
	return false
}

func converUppercaseToLowercase(char byte) byte {
	if char >= 65 && char <= 90 {
		return char + 32
	}
	return char
}

func main() {
	fmt.Println("isPalindrome is:", isPalindrome("bach"))
}
