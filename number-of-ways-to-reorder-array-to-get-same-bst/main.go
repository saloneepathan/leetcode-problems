
from typing import List
from math import comb


class Solution:
    def numOfWays(self, nums: List[int]) -> int:
        mod = 10**9 + 7

        def dfs(arr: List[int]) -> int:
            m = len(arr)

            if m < 3:
                return 1

            root = arr[0]
            left_nodes = [x for x in arr if x < root]
            right_nodes = [x for x in arr if x > root]

            left_ways = dfs(left_nodes)
            right_ways = dfs(right_nodes)

            return (
                left_ways
                * right_ways
                * comb(m - 1, len(left_nodes))
            ) % mod

        return (dfs(nums) - 1) % mod


if __name__ == "__main__":
    nums = list(map(int, input().strip().split()))
    solution = Solution()
    print(solution.numOfWays(nums))
