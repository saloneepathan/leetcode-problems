class Solution:
    def splitArray(self, nums: list[int], k: int) -> int:

        left = max(nums)
        right = sum(nums)

        while left < right:

            mid = (left + right) // 2

            groups = 1
            current_sum = 0

            for num in nums:

                if current_sum + num > mid:
                    groups += 1
                    current_sum = num
                else:
                    current_sum += num

            if groups > k:
                left = mid + 1
            else:
                right = mid

        return left

print(Solution().splitArray([7, 2, 5, 10, 8], 2))