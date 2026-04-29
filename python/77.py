from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []

        def dfs(i: int, subset: List[int]):
            if len(subset) == k:
                res.append(subset.copy())
                return

            if i > n:
                return

            # pick
            subset.append(i)
            dfs(i + 1, subset)
            subset.pop()

            # don't pick
            dfs(i + 1, subset)

        dfs(1, [])
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combine(4, 2))
    print(s.combine(1, 1))
