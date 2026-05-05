from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        res = []

        def backtrack(i: int, sol: List[str]):
            # base case
            if i == len(s):
                res.append(sol[:])
                return

            for j in range(i, len(s)):
                if self.isPalindrome(s, i, j):
                    sol.append(s[i : j + 1])
                    backtrack(j + 1, sol)
                    sol.pop()

        backtrack(0, [])
        return res

    def isPalindrome(self, s: str, i, j: int) -> bool:
        while i < j:
            if s[i] != s[j]:
                return False
            i, j = i + 1, j - 1
        return True


if __name__ == "__main__":
    s = Solution()
    print(s.partition("aab"))
    print(s.partition("a"))
