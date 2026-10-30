
import heapq
from typing import List


class Solution:
    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:

        # Edge case: at least one character is required
        if k < 1:
            return -1

        # Step 1: Build adjacency list
        graph = [[] for _ in range(n)]

        for u, v, w in edges:
            graph[u].append((v, w))

        # Step 2: Initialize distance array
        # dist[node][count] = minimum cost to reach
        # this node with count consecutive identical characters
        INF = float('inf')
        dist = [[INF] * (k + 1) for _ in range(n)]

        # Step 3: Initialize starting state
        dist[0][1] = 0

        # Min heap: (cost, node, consecutive_count)
        pq = [(0, 0, 1)]

        # Step 4: Run Dijkstra's algorithm
        while pq:

            cost, node, count = heapq.heappop(pq)

            # Ignore outdated entries
            if cost != dist[node][count]:
                continue

            # Destination reached
            if node == n - 1:
                return cost

            # Step 5: Explore neighbors
            for neighbor, weight in graph[node]:

                # Check whether character is repeated
                if labels[neighbor] == labels[node]:
                    new_count = count + 1
                else:
                    new_count = 1

                # Skip invalid paths
                if new_count > k:
                    continue

                # Calculate new path cost
                new_cost = cost + weight

                # Update if a shorter path is found
                if new_cost < dist[neighbor][new_count]:

                    dist[neighbor][new_count] = new_cost

                    heapq.heappush(
                        pq,
                        (new_cost, neighbor, new_count)
                    )

        # No valid path exists
        return -1


# -------------------------------
# Test cases
# -------------------------------

sol = Solution()

# Example 1
n = 3
edges = [[0, 1, 1], [1, 2, 1], [0, 2, 3]]
labels = "aab"
k = 1
print(sol.shortestPath(n, edges, labels, k))
# Output: 3


# Example 2
n = 3
edges = [[0, 1, 1], [1, 2, 1], [0, 2, 3]]
labels = "aab"
k = 2
print(sol.shortestPath(n, edges, labels, k))
# Output: 2


# Example 3
n = 3
edges = [[0, 1, 1], [1, 2, 1]]
labels = "aaa"
k = 2
print(sol.shortestPath(n, edges, labels, k))
# Output: -1
