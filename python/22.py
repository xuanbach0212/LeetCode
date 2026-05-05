from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def backtrack(sol: List[str], openB: int, closeB: int):
            if openB == closeB == n:
                res.append("".join(sol))
                return

            if openB < n:
                sol.append("(")
                backtrack(sol, openB + 1, closeB)
                sol.pop()
            if closeB < openB:
                sol.append(")")
                backtrack(sol, openB, closeB + 1)
                sol.pop()

        backtrack([], 0, 0)
        return res


if __name__ == "__main__":
    s = Solution()
    print("res", s.generateParenthesis(3))
    print("res", s.generateParenthesis(1))
