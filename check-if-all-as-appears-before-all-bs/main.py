class Solution:
    def checkString(self, s: str) -> bool:
        return "ba" not in s


if __name__ == "__main__":
    s = input().strip()

    solution = Solution()
    result = solution.checkString(s)

    print(result)