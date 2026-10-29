def rotate(nums: list[int], k: int) -> None:

    n = len(nums)

    if n == 0:
        return 

    k = k % n

    def reverse(left, right):
        while left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1


    reverse(0, n-1)
    reverse(0, k-1)
    reverse(k, n-1)

    print(nums)
    
print(rotate([1,2,3,4,5,6,7], 3))  # Output: [5,6,7,1,2,3,4]