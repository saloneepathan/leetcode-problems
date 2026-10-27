class Solution:
    def maxSatisfied(self, customers: list[int], grumpy: list[int], minutes: int) -> int:
        already_satisfied = 0

        for i in range(len(customers)):
            if grumpy[i] == 0:
                already_satisfied += customers[i]

        extra = 0
        for i in range(minutes):
            if grumpy[i] == 1:
                extra += customers[i]

        max_extra = extra

        for right in range(minutes, len(customers)):
            outgoing = right - minutes

            if grumpy[outgoing] == 1:
                extra -= customers[outgoing]

            if grumpy[right] == 1:
                extra += customers[right]

            max_extra = max(max_extra, extra)
        return already_satisfied + max_extra



print(Solution().maxSatisfied([1,0,1,2,1,1,7,5], [0,1,0,1,0,1,0,1], 3))  # Output: 16 