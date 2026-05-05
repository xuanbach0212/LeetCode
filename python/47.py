from typing import List


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        perm = []
        count = {n: 0 for n in nums}
        for n in nums:
            count[n] += 1

        def dfs():
            if len(nums) == len(perm):
                res.append(perm[:])
                return

            for n in count:
                if count[n] > 0:
                    perm.append(n)
                    count[n] -= 1
                    dfs()
                    count[n] += 1
                    perm.pop()

        dfs()
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.permuteUnique([1, 1, 2]))
    print(s.permuteUnique([1, 2, 3]))
