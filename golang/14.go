//go:build ignore

package main

// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
//
//
// Example 1:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:
//
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
//
// Constraints:
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters if it is non-empty

func longestCommonPrefix(strs []string) string {
	res := []byte{}

	for i := range len(strs[0]) {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return string(res)
			}
		}
		res = append(res, strs[0][i])
	}

	return string(res)
}

func main() {
}
