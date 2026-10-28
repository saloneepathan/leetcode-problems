def twoSum(nums, target):

    nums_dict = {}

    for i in range(len(nums)):

        complement = target - nums[i]

        if complement not in nums_dict:
            nums_dict[nums[i]] = i

        else: 
            return [nums_dict[complement], i]

print(twoSum([2, 7, 11, 15], 9))  # Output: [0, 1]