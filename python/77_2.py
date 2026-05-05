from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []

        def backtrack(i: int, sol: List[int]):
            if len(sol) == k:
                res.append(sol[:])
                return

            if i > n:
                return

            sol.append(i)
            backtrack(i + 1, sol)
            sol.pop()
            backtrack(i + 1, sol)

        backtrack(1, [])
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combine(4, 2))
    print(s.combine(1, 1))
