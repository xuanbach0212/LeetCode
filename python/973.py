from typing import List
import math
import heapq


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        euclide_dist = [0] * len(points)
        for i in range(len(points)):
            euclide_dist[i] = (
                math.sqrt((points[i][0] ** 2) + (points[i][1] ** 2)),
                points[i],
            )

        heapq.heapify(euclide_dist)

        res = []
        for i in range(k):
            min_dist = heapq.heappop(euclide_dist)
            res.append(min_dist[1])

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.kClosest([[1, 3], [-2, 2]], 1))
    print(s.kClosest(points=[[3, 3], [5, -1], [-2, 4]], k=2))
