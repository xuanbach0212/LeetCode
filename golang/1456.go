//go:build ignore

package main

import (
	"fmt"
)

// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
//
// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
//
//
//
// Example 1:
//
// Input: s = "abciiidef", k = 3
// Output: 3
// Explanation: The substring "iii" contains 3 vowel letters.
// Example 2:
//
// Input: s = "aeiou", k = 2
// Output: 2
// Explanation: Any substring of length 2 contains 2 vowels.
// Example 3:
//
// Input: s = "leetcode", k = 3
// Output: 2
// Explanation: "lee", "eet" and "ode" contain 2 vowels.
//
//
// Constraints:
//
// 1 <= s.length <= 105
// s consists of lowercase English letters.
// 1 <= k <= s.length

func maxVowels(s string, k int) int {
	res := 0
	sum := 0
	for i, c := range s {
		if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' {
			sum += 1
		}

		if i >= k-1 {
			res = max(sum, res)
			leftC := s[i-k+1]
			if leftC == 'a' || leftC == 'i' || leftC == 'u' || leftC == 'e' || leftC == 'o' {
				sum -= 1
			}
		}

		if sum == k {
			return k
		}
	}

	return res
}

func main() {
	// fmt.Println(totalFruit([]int{1, 2, 1}))
	fmt.Println(maxVowels("abciiidef", 3))
}
