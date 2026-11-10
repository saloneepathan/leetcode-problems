from typing import List
import random


class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        INF = float("inf")
        # dp[j] = minimum cost to reach step j; step 0 is free
        dp = [INF] * (n + 1)
        dp[0] = 0
        for j in range(1, n + 1):
            # arrive from j-1, j-2 or j-3 (jump length k)
            for k in (1, 2, 3):
                if j - k >= 0:
                    dp[j] = min(dp[j], dp[j - k] + k * k)
            dp[j] += costs[j - 1]  # costs is 1-indexed in the statement
        return dp[n]


def brute(n, costs):
    # plain recursion, no memo: only for tiny n
    def go(i):
        if i == n:
            return 0
        best = float("inf")
        for k in (1, 2, 3):
            j = i + k
            if j <= n:
                best = min(best, costs[j - 1] + k * k + go(j))
        return best
    return go(0)


if __name__ == "__main__":
    sol = Solution()
    tests = [
        (4, [1, 2, 3, 4], 13),
        (4, [5, 1, 6, 2], 11),
        (3, [9, 8, 3], 12),
        (1, [7], 8),          # edge: single step
        (2, [1, 1], 4),       # edge: 0->1->2 (2+2) vs 0->2 (1+4=5)
    ]
    ok = True
    for n, c, exp in tests:
        act = sol.climbStairs(n, c)
        ok &= act == exp
        print(f"n={n}, costs={c} | expected={exp} | actual={act} | {'PASS' if act == exp else 'FAIL'}")

    # cross-check with brute force on random small inputs
    random.seed(1)
    for _ in range(300):
        n = random.randint(1, 12)
        c = [random.randint(1, 20) for _ in range(n)]
        assert sol.climbStairs(n, c) == brute(n, c), (n, c)
    print("random cross-check vs brute force: 300 cases OK")
    print("ALL PASS" if ok else "SOME FAILED")
