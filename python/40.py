from typing import List


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        n = len(candidates)
        res = []
        candidates.sort()

        def dfs(i: int, total: int, curr: List[int]):
            if total == target:
                res.append(curr[:])
                return

            if total > target or i == len(candidates):
                return
            # pick
            curr.append(candidates[i])
            dfs(i + 1, total + candidates[i], curr)
            curr.pop()

            # don't pick
            while i + 1 < n and candidates[i] == candidates[i + 1]:
                i += 1
            dfs(i + 1, total, curr)

        dfs(0, 0, [])

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combinationSum2([10, 1, 2, 7, 6, 1, 5], 8))
    print(s.combinationSum2([2, 5, 2, 1, 2], 5))
