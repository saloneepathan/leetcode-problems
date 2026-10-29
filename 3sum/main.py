def threeSum(nums: list[int]) -> list[list[int]]:
    nums.sort()
    result = []
    n = len(nums)

    for i in range(n-2):
        if i > 0 and nums[i] == nums[i-1]:
            continue

        left = i+1
        right = n-1 

        while left < right:
            sum = nums[i] + nums[left] + nums[right]

            if sum == 0:
                result.append([nums[i], nums[left], nums[right]])

                left += 1
                right -= 1 

                while left < right and nums[left] == nums[left-1]:
                    left += 1

                while left < right and nums[right] == nums[right+1]:
                    right -= 1

            elif sum < 0:
                left += 1
            else:
                right -= 1
    
    return result
                                                      


print(threeSum([-1, 0, 1, 2, -1, -4])) 