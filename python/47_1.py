from typing import List


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        res = []
        nums.sort()
        counts = {n: 0 for n in nums}
        for n in nums:
            counts[n] += 1

        def backtrack(sol: List[int]):
            if len(sol) == len(nums):
                res.append(sol[:])
                return

            for n in counts:
                if counts[n] > 0:
                    counts[n] -= 1
                    sol.append(n)
                    backtrack(sol)
                    sol.pop()
                    counts[n] += 1

        backtrack([])

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.permuteUnique([1, 1, 2]))
    print(s.permuteUnique([1, 2, 3]))
