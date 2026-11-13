from typing import List
from collections import deque


class Solution:
    def countNonDecreasingSubarrays(self, nums: List[int], k: int) -> int:
        n = len(nums)

        # Monotonic deque.
        # Each element is [value, count].
        dq = deque()

        ans = 0
        cost = 0

        # Right boundary of the current valid window.
        right = n - 1

        # Process the left boundary from right to left.
        for left in range(n - 1, -1, -1):
            value = nums[left]
            count = 1

            # Merge all blocks smaller than the current value.
            while dq and dq[-1][0] < value:
                smaller_value, smaller_count = dq.pop()

                cost += (value - smaller_value) * smaller_count
                count += smaller_count

            dq.append([value, count])

            # If the cost is too large, move the right boundary left.
            while cost > k:
                value_at_right = nums[right]

                # Remove this element's contribution.
                cost -= dq[0][0] - value_at_right

                right -= 1

                # Remove one element from the front block.
                if dq[0][1] == 1:
                    dq.popleft()
                else:
                    dq[0][1] -= 1

            # All subarrays:
            #
            # [left, left]
            # [left, left + 1]
            # ...
            # [left, right]
            #
            # are valid.
            ans += right - left + 1

        return ans


def main():
    solution = Solution()

    # Example 1
    nums = [6, 3, 1, 2, 4, 4]
    k = 7

    result = solution.countNonDecreasingSubarrays(nums, k)

    print("Example 1:")
    print("nums =", nums)
    print("k =", k)
    print("Answer =", result)
    print("Expected = 17")
    print()

    # Example 2
    nums = [6, 3, 1, 3, 6]
    k = 4

    result = solution.countNonDecreasingSubarrays(nums, k)

    print("Example 2:")
    print("nums =", nums)
    print("k =", k)
    print("Answer =", result)
    print("Expected = 12")


if __name__ == "__main__":
    main()