
from typing import List
from collections import defaultdict
from heapq import heappush, heappop
from math import inf


class Solution:
    def minOperations(self, nums: List[int], x: int, k: int) -> int:
        n = len(nums)

        left, right = [], []
        delayed = defaultdict(int)

        left_size = right_size = 0
        sum_l = sum_r = 0

        ops = [0] * n

        def prune(h: List[int], sign: int) -> None:
            while h and delayed[sign * h[0]]:
                delayed[sign * h[0]] -= 1
                heappop(h)

        def balance() -> None:
            nonlocal left_size, right_size, sum_l, sum_r

            if left_size > right_size + 1:
                num = -heappop(left)

                heappush(right, num)

                sum_r += num
                sum_l -= num

                right_size += 1
                left_size -= 1

                prune(left, -1)

            elif right_size > left_size:
                num = heappop(right)

                heappush(left, -num)

                sum_l += num
                sum_r -= num

                left_size += 1
                right_size -= 1

                prune(right, 1)

        # Calculate minimum operations for every window of size x.
        for i, num in enumerate(nums):

            if not left or num <= -left[0]:
                heappush(left, -num)
                left_size += 1
                sum_l += num
            else:
                heappush(right, num)
                right_size += 1
                sum_r += num

            balance()

            # Remove the element that leaves the sliding window.
            if i >= x:
                num = nums[i - x]

                delayed[num] += 1

                if num <= -left[0]:
                    left_size -= 1
                    sum_l -= num

                    if num == -left[0]:
                        prune(left, -1)

                else:
                    right_size -= 1
                    sum_r -= num

                    if num == right[0]:
                        prune(right, 1)

                balance()

            # Calculate the cost of making this window equal.
            if i >= x - 1:
                mid = -left[0]

                ops[i] = (
                    sum_r - sum_l
                    + mid * (left_size - right_size)
                )

        # DP initialization.
        dp = [[0] + [inf] * k for _ in range(n)]

        for j in range(1, k + 1):
            if j == 1:
                dp[x - 1][j] = ops[x - 1]
            else:
                dp[j * x - 1][j] = (
                    dp[(j - 1) * x - 1][j - 1]
                    + ops[j * x - 1]
                )

        # DP transitions.
        for j in range(1, k + 1):
            for i in range(j * x, n):
                dp[i][j] = min(
                    dp[i - x][j - 1] + ops[i],
                    dp[i - 1][j]
                )

        return dp[n - 1][k]


# Example usage
if __name__ == "__main__":
    solution = Solution()

    nums = [2, 3, 1, 4, 5]
    x = 3
    k = 1

    result = solution.minOperations(nums, x, k)

    print("Minimum operations:", result)
