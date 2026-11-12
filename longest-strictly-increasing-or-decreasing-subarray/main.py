from typing import List


class Solution:
    def longestMonotonicSubarray(self, nums: List[int]) -> int:
        inc = 1
        dec = 1
        ans = 1

        for i in range(1, len(nums)):
            if nums[i] > nums[i - 1]:
                inc += 1
                dec = 1

            elif nums[i] < nums[i - 1]:
                dec += 1
                inc = 1

            else:
                inc = 1
                dec = 1

            ans = max(ans, inc, dec)

        return ans


# Example usage
if __name__ == "__main__":
    solution = Solution()

    nums = [1, 4, 3, 3, 2]

    result = solution.longestMonotonicSubarray(nums)

    print("Input:", nums)
    print("Output:", result)