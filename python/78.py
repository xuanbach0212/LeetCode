from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        res = []
        subset = []

        def dfs(i):
            if i == n:
                res.append(subset[:])
                return

            # don pick
            dfs(i + 1)
            # pick
            subset.append(nums[i])
            dfs(i + 1)
            subset.pop()

        dfs(0)
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.subsets([1, 2, 3]))
    print(s.subsets([0]))
