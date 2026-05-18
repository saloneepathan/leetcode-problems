from typing import List
from collections import defaultdict

class Solution:
    def minJumps(self, arr: List[int]) -> int:
        n = len(arr)

        # If array has only one element
        if n <= 1:
            return 0

        # Store indices for each value
        graph = defaultdict(list)
        for i, value in enumerate(arr):
            graph[value].append(i)

        # BFS initialization
        curs = [0]
        visited = {0}
        step = 0

        while curs:
            nex = []

            for node in curs:

                # Reached last index
                if node == n - 1:
                    return step

                # Jump to same value indices
                for child in graph[arr[node]]:
                    if child not in visited:
                        visited.add(child)
                        nex.append(child)

                # Clear to avoid repeated processing
                graph[arr[node]].clear()

                # Jump to adjacent indices
                for child in [node - 1, node + 1]:
                    if 0 <= child < n and child not in visited:
                        visited.add(child)
                        nex.append(child)

            curs = nex
            step += 1

        return -1


# Example usage
if __name__ == "__main__":
    arr = [100, -23, -23, 404, 100, 23, 23, 23, 3, 404]

    sol = Solution()
    result = sol.minJumps(arr)

    print("Minimum jumps:", result)