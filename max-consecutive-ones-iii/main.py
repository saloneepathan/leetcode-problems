def longestOnes(nums: list[int], k: int) -> int:
    left = 0
    zeroes = 0
    longest = 0

    for right in range(len(nums)):
        if nums[right] == 0:
            zeroes += 1

        while zeroes > k:
            if nums[left] == 0:
                zeroes -= 1

            left += 1
        longest = max(longest, right - left + 1)

    return longest

print("Length of longest substring with at most k zeroes:", longestOnes([1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 3))
