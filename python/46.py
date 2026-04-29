from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res, sol = [], []

        def backtrack():
            if len(nums) == len(sol):
                res.append(sol[:])
                return

            for num in nums:
                if num not in sol:
                    sol.append(num)
                    backtrack()
                    sol.pop()

        backtrack()
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.permute([1, 2, 3]))
    print(s.permute([0, 1]))
