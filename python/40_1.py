from typing import List


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        res = []
        candidates.sort()

        def backtrack(i: int, sol: List[int], total: int):
            # base case
            if total == target:
                res.append(sol[:])
                return

            if i == len(candidates) or total > target:
                return

            # pick
            sol.append(candidates[i])
            backtrack(i + 1, sol, total + candidates[i])
            sol.pop()

            # don't pick
            while i + 1 < len(candidates) and candidates[i] == candidates[i + 1]:
                i = i + 1
            backtrack(i + 1, sol, total)

        backtrack(0, [], 0)
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combinationSum2([10, 1, 2, 7, 6, 1, 5], 8))
    print(s.combinationSum2([2, 5, 2, 1, 2], 5))
