
class Solution:
    def minCost(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])

        dp = [[set() for _ in range(n)] for _ in range(m)]

        dp[0][0].add(grid[0][0])

        for i in range(m):
            for j in range(n):
                if i == 0 and j == 0:
                    continue

                value = grid[i][j]

                if i > 0:
                    for xor_val in dp[i - 1][j]:
                        dp[i][j].add(xor_val ^ value)

                if j > 0:
                    for xor_val in dp[i][j - 1]:
                        dp[i][j].add(xor_val ^ value)

        return min(dp[m - 1][n - 1])


# Run test cases
if __name__ == "__main__":
    solution = Solution()

    grid1 = [[1, 2], [3, 4]]
    print(solution.minCost(grid1))  # Output: 6

    grid2 = [[6, 7], [5, 8]]
    print(solution.minCost(grid2))  # Output: 9

    grid3 = [[2, 7, 5]]
    print(solution.minCost(grid3))  # Output: 0
