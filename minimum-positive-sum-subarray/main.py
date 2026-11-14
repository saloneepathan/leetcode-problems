from typing import List
class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        n = len(nums)
        ans = float("inf")
        for i in range(n):
            current_sum = 0
            for j in range(i, min(n, i + r)):
                current_sum += nums[j]
                length = j - i + 1
                if length >= l and current_sum > 0:
                    ans = min(ans, current_sum)
        return -1 if ans == float("inf") else ans
# -------------------------
# Test
# -------------------------
nums = [3, -2, 1, 4]
l = 2
r = 3
solution = Solution()
result = solution.minimumSumSubarray(nums, l, r)
print(result)

