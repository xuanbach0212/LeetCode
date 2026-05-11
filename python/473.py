from typing import List


class Solution:
    def makesquare(self, matchsticks: List[int]) -> bool:
        def backtrack(i: int, sol: List[int]) -> bool:
            if len(sol) == 4:
                if sol[0] == sol[1] == sol[2] == sol[3]:
                    return True
                return False

            if i == len(matchsticks):
                return False

            res = False
            sol.append(matchsticks[i])
            res = backtrack(i + 1, sol)
            sol.pop()
            return res

        return backtrack(0, [])


if __name__ == "__main__":
    s = Solution()
    print(s.makesquare([1, 1, 2, 2, 2]))
    print(s.makesquare([3, 3, 3, 3, 4]))
