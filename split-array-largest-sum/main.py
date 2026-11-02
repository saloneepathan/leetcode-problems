def splitArray(nums, k):
    left = max(nums)
    right = sum(nums)

    while left < right:
        mid = (left + right) // 2


        groups = 1
        currSum = 0

        for num in nums:
            if currSum + num > mid:
                groups += 1
                currSum = num 
            else:
                currSum += num

        if groups > k:
            left = left + 1
        else: 
            right = mid 

    return left

print(splitArray([7, 3, 4, 8, 10, 9, 4], 2))