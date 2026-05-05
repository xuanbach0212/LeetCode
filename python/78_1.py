from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []

        def backtrack(i: int, subset: List[int]):
            # base case
            if i == len(nums):
                res.append(subset[:])
                return

            # pick
            subset.append(nums[i])
            backtrack(i + 1, subset)
            subset.pop()
            backtrack(i + 1, subset)

        backtrack(0, [])
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.subsets([1, 2, 3]))
    print(s.subsets([0]))
