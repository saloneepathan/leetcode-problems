from functools import cache
from math import isqrt


def segmentCost(s: str, l: int, m: int, factors: list[list[int]]) -> int:
    best = m

    for d in factors[m]:
        groups = m // d
        changes = 0

        for offset in range(d):
            for t in range(groups // 2):
                left = l + offset + t * d
                right = l + offset + (groups - 1 - t) * d

                if s[left] != s[right]:
                    changes += 1

        best = min(best, changes)

    return best


class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)

        # Precompute proper divisors for every possible substring length.
        factors = [[] for _ in range(n + 1)]

        for d in range(1, n // 2 + 1):
            for m in range(d * 2, n + 1, d):
                factors[m].append(d)

        # cost[l][r] = minimum changes needed to make
        # s[l:r+1] a palindrome-like valid segment.
        cost = [[0] * n for _ in range(n)]

        for l in range(n - 1):
            for r in range(l + 1, n):
                length = r - l + 1
                cost[l][r] = segmentCost(
                    s, l, length, factors
                )

        @cache
        def solve(seg: int, start: int) -> int:
            # Only one segment remains.
            if seg == 1:
                return cost[start][n - 1]

            # Leave enough characters for the remaining segments.
            limit = n - 1 - 2 * (seg - 1)

            return min(
                cost[start][end] + solve(seg - 1, end + 1)
                for end in range(start + 1, limit + 1)
            )

        return solve(k, 0)


# ---------------------------------------------------------
# Runnable test
# ---------------------------------------------------------

if __name__ == "__main__":
    solution = Solution()

    test_cases = [
        # Add today's LeetCode examples here if needed.
    ]

    for s, k, expected in test_cases:
        result = solution.minimumChanges(s, k)
        print(f"s={s!r}, k={k}")
        print(f"Output:   {result}")
        print(f"Expected: {expected}")
        print(f"Pass:     {result == expected}")
        print()