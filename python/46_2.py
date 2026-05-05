from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        used = [False] * len(nums)

        def backtrack(sol: List[int]):
            if len(sol) == len(nums):
                res.append(sol[:])
                return

            for i in range(len(nums)):
                if not used[i]:
                    used[i] = True
                    sol.append(nums[i])
                    backtrack(sol)
                    sol.pop()
                    used[i] = False

        backtrack([])
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.permute([1, 2, 3]))
    print(s.permute([0, 1]))
