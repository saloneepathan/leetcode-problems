
from collections import deque

class Solution:
    def maxSlidingWindow(self, nums: list[int], k: int) -> list[int]:
        dq = deque()
        ans = []

        for right in range(len(nums)):

            # 1. Remove indices outside the window
            while dq and dq[0] <= right - k:
                dq.popleft()

            # 2. Remove smaller or equal values
            while dq and nums[dq[-1]] <= nums[right]:
                dq.pop()

            # 3. Add the current index
            dq.append(right)

            # 4. Record maximum once window is full
            if right >= k - 1:
                ans.append(nums[dq[0]])

        return ans

print(Solution().maxSlidingWindow([1,3,-1,-3,5,3,6,7], 3))  # Output: [3,3,5,5,6,7]