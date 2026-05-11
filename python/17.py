from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        telephone = ["abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"]
        res = []

        def backtrack(i: int, sol: List[str]):
            if len(sol) == len(digits):
                res.append("".join(sol))
                return

            idxTele = int(digits[i]) - 2
            for c in telephone[idxTele]:
                sol.append(c)
                backtrack(i + 1, sol)
                sol.pop()

        backtrack(0, [])

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.letterCombinations("23"))
    print(s.letterCombinations("2"))
