
from typing import List


class Solution:
    def palindromicPathQueries(
        self,
        n: int,
        edges: List[List[int]],
        s: str,
        queries: List[str]
    ) -> List[bool]:

        # Build adjacency list
        adj = [[] for _ in range(n)]
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)

        def char_mask(c):
            return 1 << (ord(c) - ord('a'))

        chars = list(s)

        # Euler Tour, depth, parent, and initial prefix masks
        tin = [0] * n
        tout = [0] * n
        depth = [0] * n
        parent = [0] * n
        prefix = [0] * n

        timer = 0
        stack = [(0, -1, False)]

        while stack:
            u, p, exiting = stack.pop()

            if exiting:
                tout[u] = timer - 1
                continue

            parent[u] = u if p == -1 else p
            tin[u] = timer
            timer += 1

            if p == -1:
                prefix[u] = char_mask(chars[u])
            else:
                depth[u] = depth[p] + 1
                prefix[u] = prefix[p] ^ char_mask(chars[u])

            stack.append((u, p, True))

            for v in reversed(adj[u]):
                if v != p:
                    stack.append((v, u, False))

        # Binary lifting table for LCA
        LOG = n.bit_length()
        up = [parent[:]]

        for k in range(1, LOG):
            prev = up[k - 1]
            up.append([prev[prev[u]] for u in range(n)])

        def lca(u, v):
            if depth[u] < depth[v]:
                u, v = v, u

            diff = depth[u] - depth[v]

            for k in range(LOG):
                if diff & (1 << k):
                    u = up[k][u]

            if u == v:
                return u

            for k in range(LOG - 1, -1, -1):
                if up[k][u] != up[k][v]:
                    u = up[k][u]
                    v = up[k][v]

            return parent[u]

        # Fenwick Tree supporting range XOR updates
        # and point queries
        bit = [0] * (n + 2)

        def add(index, delta):
            while index < len(bit):
                bit[index] ^= delta
                index += index & -index

        def point_query(index):
            result = 0

            while index > 0:
                result ^= bit[index]
                index -= index & -index

            return result

        def range_xor(left, right, delta):
            add(left + 1, delta)
            add(right + 2, delta)

        def current_prefix(u):
            return prefix[u] ^ point_query(tin[u] + 1)

        # Process queries
        answer = []

        for q in queries:
            parts = q.split()

            if parts[0] == "update":
                u = int(parts[1])
                new_char = parts[2]

                delta = char_mask(chars[u]) ^ char_mask(new_char)

                range_xor(tin[u], tout[u], delta)
                chars[u] = new_char

            else:
                u = int(parts[1])
                v = int(parts[2])

                w = lca(u, v)

                mask = (
                    current_prefix(u)
                    ^ current_prefix(v)
                    ^ char_mask(chars[w])
                )

                answer.append((mask & (mask - 1)) == 0)

        return answer


# -------------------------
# Run example test cases
# -------------------------
if __name__ == "__main__":
    solution = Solution()

    # Example 1
    n = 3
    edges = [[0, 1], [1, 2]]
    s = "aac"
    queries = [
        "query 0 2",
        "update 1 b",
        "query 0 2"
    ]

    print(solution.palindromicPathQueries(n, edges, s, queries))
    # Expected: [True, False]

    # Example 2
    n = 4
    edges = [[0, 1], [0, 2], [0, 3]]
    s = "abca"
    queries = [
        "query 1 2",
        "update 0 b",
        "query 2 3",
        "update 3 a",
        "query 1 3"
    ]

    print(solution.palindromicPathQueries(n, edges, s, queries))
    # Expected: [False, False, True]
