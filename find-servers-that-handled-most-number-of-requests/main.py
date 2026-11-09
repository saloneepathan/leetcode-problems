import heapq
from typing import List


class Solution:
    def busiestServers(self, k: int, arrival: List[int], load: List[int]) -> List[int]:
        count = [0] * k
        busy = []  # (end_time, server)
        # Free servers stored as "virtual indices" so that, for request i,
        # the smallest virtual index >= i maps to the next server at or after i % k.
        available = list(range(k))

        for i, (start, duration) in enumerate(zip(arrival, load)):
            # Release servers that are done by this arrival time.
            while busy and busy[0][0] <= start:
                _, server = heapq.heappop(busy)
                # Re-key the server so it sorts relative to request i:
                # i + ((server - i) % k) is the smallest value >= i that is ≡ server (mod k).
                heapq.heappush(available, i + (server - i) % k)

            if not available:
                continue  # request dropped

            # Invariant: every key in `available` is >= i (each request pops the min,
            # and keys are distinct mod k), so the min key is the right server.
            server = heapq.heappop(available) % k
            count[server] += 1
            heapq.heappush(busy, (start + duration, server))

        best = max(count)
        return [s for s, c in enumerate(count) if c == best]


if __name__ == "__main__":
    s = Solution()
    print(s.busiestServers(3, [1, 2, 3, 4, 5], [5, 2, 3, 3, 3]))  # [1]
    print(s.busiestServers(3, [1, 2, 3, 4], [1, 2, 1, 2]))  # [0]
    print(s.busiestServers(3, [1, 2, 3], [10, 12, 11]))  # [0, 1, 2]
