def increasingSubsequence(nums: list[int]) -> bool:
    first = second = float("inf")

    for i in nums:
        if i < first:
            first = i
        elif i < second:
            second = i
        else:
            return True
    return False

print(increasingSubsequence([5, 4, 3, 0])) 
    
