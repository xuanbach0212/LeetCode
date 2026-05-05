from typing import List


class Solution:
    def subsetXORSum(self, nums: List[int]) -> int:
        def backtrack(i: int, total: int) -> int:
            if i == len(nums):
                return total

            return backtrack(i + 1, total) + backtrack(i + 1, total ^ nums[i])

        return backtrack(0, 0)


if __name__ == "__main__":
    s = Solution()
    print(s.subsetXORSum([1, 3]))
    print(s.subsetXORSum([5, 1, 6]))
