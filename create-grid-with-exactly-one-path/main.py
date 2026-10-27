from typing import List

class Solution:
    def constructGrid(self, m: int, n: int) -> List[str]:
        grid = [['#'] * n for _ in range(m)]

        # Make the first row free
        for j in range(n):
            grid[0][j] = '.'

        # Make the last column free
        for i in range(m):
            grid[i][n - 1] = '.'

        # Convert each row into a string
        return [''.join(row) for row in grid]


# Run the code
if __name__ == "__main__":
    sol = Solution()

    print(sol.constructGrid(2, 3))
    print(sol.constructGrid(3, 3))
    print(sol.constructGrid(1, 4))