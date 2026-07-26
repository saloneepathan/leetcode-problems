from typing import List

class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        left, right = 0, len(letters) - 1
        ans = letters[0]  # default for wrap-around case

        while left <= right:
            mid = (left + right) // 2
            if letters[mid] > target:
                ans = letters[mid]
                right = mid - 1
            else:
                left = mid + 1

        return ans


# --------- Driver Code (for testing) ---------
if __name__ == "__main__":
    solution = Solution()

    print(solution.nextGreatestLetter(["c", "f", "j"], "a"))  # Output: c
    print(solution.nextGreatestLetter(["c", "f", "j"], "c"))  # Output: f
    print(solution.nextGreatestLetter(["x", "x", "y", "y"], "z"))  # Output: x
