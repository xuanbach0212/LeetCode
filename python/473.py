from typing import List


class Solution:
    def makesquare(self, matchsticks: List[int]) -> bool:
        length = sum(matchsticks) // 4
        if sum(matchsticks) % 4 != 0:
            return False

        matchsticks.sort(reverse=True)
        sides = [0] * 4

        def backtrack(i: int) -> bool:
            if i == len(matchsticks):
                return True

            for j in range(4):
                if sides[j] + matchsticks[i] <= length:
                    sides[j] += matchsticks[i]
                    if backtrack(i + 1):
                        return True
                    sides[j] -= matchsticks[i]

            return False

        return backtrack(0)


if __name__ == "__main__":
    s = Solution()
    print(s.makesquare([1, 1, 2, 2, 2]))
    print(s.makesquare([3, 3, 3, 3, 4]))
