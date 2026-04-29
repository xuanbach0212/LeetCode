from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        n = len(candidates)
        res = []

        def dfs(i: int, curr: List[int], total: int):
            if total == target:
                res.append(curr.copy())
                return

            if total > target or i >= n:
                return

            # don't pick
            dfs(i + 1, curr, total)

            # pick
            curr.append(candidates[i])
            dfs(i, curr, total + candidates[i])
            curr.pop()

        dfs(0, [], 0)
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combinationSum([2, 3, 6, 7], 7))
    print(s.combinationSum([2, 3, 5], 8))
