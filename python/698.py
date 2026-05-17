from typing import List


class Solution:
    def canPartitionKSubsets(self, nums: List[int], k: int) -> bool:
        if sum(nums) % k != 0:
            return False
        target = sum(nums) // k
        used = [False] * len(nums)
        nums.sort(reverse=True)

        def backtrack(i: int, sol: int, k: int):
            if k == 0:
                return True

            if sol == target:
                return backtrack(0, 0, k - 1)

            for j in range(i, len(nums)):
                if used[j] or sol + nums[j] > target:
                    continue

                used[j] = True
                if backtrack(j + 1, sol + nums[j], k):
                    return True
                used[j] = False

                if sol == 0:
                    return False

            return False

        return backtrack(0, 0, k)


if __name__ == "__main__":
    s = Solution()
    print(s.canPartitionKSubsets(nums=[4, 3, 2, 3, 5, 2, 1], k=4))
    print(s.canPartitionKSubsets(nums=[1, 2, 3, 4], k=3))
    print(s.canPartitionKSubsets(nums=[2, 2, 2, 2, 3, 4, 5], k=4))
    print(
        s.canPartitionKSubsets(
            [10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8], k=11
        )
    )
