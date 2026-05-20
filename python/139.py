from typing import List

# Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
#
# Note that the same word in the dictionary may be reused multiple times in the segmentation.
#
#
#
# Example 1:
#
# Input: s = "leetcode", wordDict = ["leet","code"]
# Output: true
# Explanation: Return true because "leetcode" can be segmented as "leet code".
#
# Example 2:
#
# Input: s = "applepenapple", wordDict = ["apple","pen"]
# Output: true
# Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
# Note that you are allowed to reuse a dictionary word.
#
# Example 3:
#
# Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
# Output: false
#
#
#
# Constraints:
#
# 1 <= s.length <= 300
# 1 <= wordDict.length <= 1000
# 1 <= wordDict[i].length <= 20
# s and wordDict[i] consist of only lowercase English letters.
# All the strings of wordDict are unique.


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        visited = set()

        def backtrack(i: int) -> bool:
            if i == len(s):
                return True

            if i in visited:
                return False

            for word in wordDict:
                if i + len(word) > len(s) or s[i : i + len(word)] != word:
                    continue
                if backtrack(i + len(word)):
                    return True
            visited.add(i)
            return False

        return backtrack(0)


if __name__ == "__main__":
    s = Solution()
    print(s.wordBreak(s="leetcode", wordDict=["leet", "code"]))
    print(s.wordBreak(s="applepenapple", wordDict=["apple", "pen"]))
    print(s.wordBreak(s="catsandog", wordDict=["cats", "dog", "sand", "and", "cat"]))
    print(s.wordBreak(s="cars", wordDict=["car", "ca", "rs"]))
    print(
        s.wordBreak(
            s="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
            wordDict=[
                "a",
                "aa",
                "aaa",
                "aaaa",
                "aaaaa",
                "aaaaaa",
                "aaaaaaa",
                "aaaaaaaa",
                "aaaaaaaaa",
                "aaaaaaaaaa",
            ],
        )
    )
