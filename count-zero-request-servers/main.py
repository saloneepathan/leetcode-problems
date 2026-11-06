from typing import List


class Solution:
    def countServers(
        self,
        n: int,
        logs: List[List[int]],
        x: int,
        queries: List[int]
    ) -> List[int]:

        # Sort logs by time
        logs.sort(key=lambda log: log[1])

        # Sort queries while keeping their original indices
        sorted_queries = sorted(
            (query, i) for i, query in enumerate(queries)
        )

        result = [0] * len(queries)

        # Frequency of each server in the current window
        freq = [0] * (n + 1)

        left = 0
        right = 0

        # Number of distinct servers that received a request
        active_servers = 0

        for query, original_index in sorted_queries:

            start = query - x
            end = query

            # Add logs whose time <= query
            while right < len(logs) and logs[right][1] <= end:
                server_id = logs[right][0]

                if freq[server_id] == 0:
                    active_servers += 1

                freq[server_id] += 1
                right += 1

            # Remove logs whose time < query - x
            # Interval is inclusive, so logs at query - x stay
            while left < right and logs[left][1] < start:
                server_id = logs[left][0]

                freq[server_id] -= 1

                if freq[server_id] == 0:
                    active_servers -= 1

                left += 1

            # Servers with no requests in [query - x, query]
            result[original_index] = n - active_servers

        return result


# -------------------------------
# Test
# -------------------------------

if __name__ == "__main__":
    solution = Solution()

    n = 3
    logs = [[1, 3], [2, 6], [1, 5]]
    x = 5
    queries = [10, 11]

    answer = solution.countServers(n, logs, x, queries)

    print("Output:", answer)