
class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        smaller = sum(num < target for num in nums)
        equal = sum(num == target for num in nums)

        return list(range(smaller, smaller + equal))


if __name__ == "__main__":
    sol = Solution()

    nums = [1, 2, 5, 2, 3]
    target = 2
    print(sol.targetIndices(nums, target))  # Output: [1, 2]

    nums = [1, 2, 5, 2, 3]
    target = 3
    print(sol.targetIndices(nums, target))  # Output: [3]

    nums = [1, 2, 5, 2, 3]
    target = 5
    print(sol.targetIndices(nums, target))  # Output: [4]
