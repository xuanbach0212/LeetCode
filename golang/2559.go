//go:build ignore

package main

import "fmt"

// You are given a 0-indexed array of strings words and a 2D array of integers queries.
//
// Each query queries[i] = [li, ri] asks us to find the number of strings present at the indices ranging from li to ri (both inclusive) of words that start and end with a vowel.
//
// Return an array ans of size queries.length, where ans[i] is the answer to the ith query.
//
// Note that the vowel letters are 'a', 'e', 'i', 'o', and 'u'.
//
//
//
// Example 1:
//
// Input: words = ["aba","bcb","ece","aa","e"], queries = [[0,2],[1,4],[1,1]]
// Output: [2,3,0]
// Explanation: The strings starting and ending with a vowel are "aba", "ece", "aa" and "e".
// The answer to the query [0,2] is 2 (strings "aba" and "ece").
// to query [1,4] is 3 (strings "ece", "aa", "e").
// to query [1,1] is 0.
// We return [2,3,0].
// Example 2:
//
// Input: words = ["a","e","i"], queries = [[0,2],[0,1],[2,2]]
// Output: [3,2,1]
// Explanation: Every string satisfies the conditions, so we return [3,2,1].
//
//
// Constraints:
//
// 1 <= words.length <= 105
// 1 <= words[i].length <= 40
// words[i] consists only of lowercase English letters.
// sum(words[i].length) <= 3 * 105
// 1 <= queries.length <= 105
// 0 <= li <= ri < words.length

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func isValid(s string) bool {
	return isVowel(s[0]) && isVowel(s[len(s)-1])
}

func vowelStrings(words []string, queries [][]int) []int {
	prefix := make([]int, len(words)+1)
	sum := 0
	for i := 0; i < len(words); i++ {
		if isValid(words[i]) {
			sum += 1
		}
		prefix[i+1] = sum
	}

	res := []int{}
	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]
		res = append(res, prefix[r+1]-prefix[l])
	}
	return res
}

func main() {
	fmt.Println(vowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))
	fmt.Println(vowelStrings([]string{"a", "e", "i"}, [][]int{{0, 2}, {0, 1}, {2, 2}}))
}
