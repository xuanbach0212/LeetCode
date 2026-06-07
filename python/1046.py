import heapq
from typing import List


class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        for i in range(len(stones)):
            stones[i] = -stones[i]

        heapq.heapify(stones)

        largest = 0
        secondLargest = 0
        while len(stones) > 1:
            largest = -heapq.heappop(stones)
            secondLargest = -heapq.heappop(stones)

            sub = largest - secondLargest
            if sub != 0:
                heapq.heappush(stones, -sub)

        return -stones[0] if stones else 0


if __name__ == "__main__":
    s = Solution()
    print(s.lastStoneWeight(stones=[2, 7, 4, 1, 8, 1]))
    print(s.lastStoneWeight(stones=[1]))
