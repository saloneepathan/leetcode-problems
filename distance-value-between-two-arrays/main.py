from bisect import bisect_left


class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        arr2.sort()

        count = 0

        for x in arr1:
            # Find the first element >= x - d
            i = bisect_left(arr2, x - d)

            # If no element exists in [x-d, x+d], count x
            if i == len(arr2) or arr2[i] > x + d:
                count += 1

        return count


# Example 1
arr1 = [4, 5, 8]
arr2 = [10, 9, 1, 8]
d = 2

solution = Solution()

answer = solution.findTheDistanceValue(arr1, arr2, d)

print("Output:", answer)