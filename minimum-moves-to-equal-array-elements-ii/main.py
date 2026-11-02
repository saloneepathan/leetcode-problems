def minMoves(nums: list[int]) -> int:
    nums.sort()

    median = nums[len(nums)//2]

    moves = 0

    for num in nums:
        moves += abs(num-median)

    return moves

print(minMoves([1, 4, 6, 2, 3, 9, 6]))