class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        # If type 3 is not cheaper than buying type 1 and type 2 separately
        if costBoth >= cost1 + cost2:
            return cost1 * need1 + cost2 * need2

        # Use type 3 items for the common requirement
        common = min(need1, need2)

        cost = common * costBoth

        # Remaining type 1 requirement
        remaining1 = need1 - common
        cost += remaining1 * min(cost1, costBoth)

        # Remaining type 2 requirement
        remaining2 = need2 - common
        cost += remaining2 * min(cost2, costBoth)

        return cost


# Driver code
if __name__ == "__main__":
    cost1, cost2, costBoth, need1, need2 = map(int, input().split())

    solution = Solution()
    result = solution.minimumCost(
        cost1, cost2, costBoth, need1, need2
    )

    print(result)