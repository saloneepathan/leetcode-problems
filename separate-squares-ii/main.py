from typing import List
import bisect


class SegmentTree:
    def __init__(self, xs: List[int]):
        self.xs = xs
        self.n = len(xs) - 1
        self.count = [0] * (4 * self.n)
        self.covered = [0] * (4 * self.n)

    def update(self, ql, qr, val, l, r, idx):
        if self.xs[r + 1] <= ql or self.xs[l] >= qr:
            return
        if ql <= self.xs[l] and self.xs[r + 1] <= qr:
            self.count[idx] += val
        else:
            mid = (l + r) // 2
            self.update(ql, qr, val, l, mid, idx * 2 + 1)
            self.update(ql, qr, val, mid + 1, r, idx * 2 + 2)

        if self.count[idx] > 0:
            self.covered[idx] = self.xs[r + 1] - self.xs[l]
        else:
            if l == r:
                self.covered[idx] = 0
            else:
                self.covered[idx] = (
                    self.covered[idx * 2 + 1] + self.covered[idx * 2 + 2]
                )

    def query(self):
        return self.covered[0]


class Solution:
    def separateSquares(self, squares: List[List[int]]) -> float:
        events = []
        xs = set()

        # Build sweep events
        for x, y, l in squares:
            events.append((y, 1, x, x + l))
            events.append((y + l, -1, x, x + l))
            xs.add(x)
            xs.add(x + l)

        xs = sorted(xs)
        events.sort()

        seg = SegmentTree(xs)

        psum = []     # prefix areas
        widths = []   # active x-length after each event
        ys = []       # y-coordinate of each event

        total_area = 0.0
        prev_y = events[0][0]

        # Sweep line on y-axis
        for y, delta, xl, xr in events:
            cur_width = seg.query()
            total_area += cur_width * (y - prev_y)

            seg.update(xl, xr, delta, 0, seg.n - 1, 0)

            psum.append(total_area)
            widths.append(seg.query())
            ys.append(y)

            prev_y = y

        target = total_area / 2.0

        # Find segment containing target area
        i = bisect.bisect_left(psum, target)

        if i == 0:
            return ys[0]

        area_before = psum[i - 1]
        remaining = target - area_before
        width = widths[i - 1]

        # height increment needed inside this segment
        return ys[i - 1] + remaining / width
# Example usage:
sol = Solution()
print(sol.separateSquares([[0,0,2],[1,1,2]]))  # Expected output: 2.0 