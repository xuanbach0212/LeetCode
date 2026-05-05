from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res = []

        def backtrack(i: int, sol: List[int], total: int):
            # base case
            if total == target:
                res.append(sol[:])
                return

            if i == len(candidates) or total > target:
                return

            # pick
            sol.append(candidates[i])
            backtrack(i, sol, total + candidates[i])
            sol.pop()

            # don't pick
            backtrack(i + 1, sol, total)

        backtrack(0, [], 0)
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combinationSum([2, 3, 6, 7], 7))
    print(s.combinationSum([2, 3, 5], 8))
