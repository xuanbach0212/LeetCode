from typing import List


class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()

        def dfs(i: int, sol: List[int]):
            if i == len(nums):
                res.append(sol[:])
                return

            sol.append(nums[i])
            dfs(i + 1, sol)
            sol.pop()

            while i + 1 < len(nums) and nums[i] == nums[i + 1]:
                i += 1
            dfs(i + 1, sol)

        dfs(0, [])
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.subsetsWithDup([1, 2, 2]))
    print(s.subsetsWithDup([0]))
