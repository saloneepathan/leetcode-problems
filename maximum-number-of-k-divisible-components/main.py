from typing import List
import sys

class Solution:
    def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
        sys.setrecursionlimit(10**7)

        # Build adjacency list (fastest way)
        g = [[] for _ in range(n)]
        for a, b in edges:
            g[a].append(b)
            g[b].append(a)

        components = 0
        vals = values   # local bind = faster
        kk = k

        def dfs(node, parent):
            nonlocal components
            subtotal = vals[node]

            for nei in g[node]:
                if nei != parent:
                    subtotal += dfs(nei, node)

            if subtotal % kk == 0:
                components += 1
                return 0  # cut here

            return subtotal

        dfs(0, -1)
        return components
