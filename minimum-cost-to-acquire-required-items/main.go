package main

import "fmt"

// minimumCost tries buying k "both" items for k in {0, need1, need2}.
// Total cost f(k) = k*costBoth + max(0,need1-k)*cost1 + max(0,need2-k)*cost2
// is convex and piecewise linear, so its minimum is at a breakpoint.
func minimumCost(cost1, cost2, costBoth, need1, need2 int64) int64 {
	f := func(k int64) int64 {
		return k*costBoth + max(0, need1-k)*cost1 + max(0, need2-k)*cost2
	}
	return min(f(0), f(need1), f(need2))
}

func main() {
	tests := []struct {
		c1, c2, cb, n1, n2, want int64
	}{
		{3, 2, 1, 3, 2, 3},
		{5, 4, 15, 2, 3, 22},
		{5, 4, 15, 0, 0, 0},
		{5, 4, 6, 3, 1, 16}, // mixed: 1 both (6) + 2 type1 (10)
		{1000000, 1000000, 1, 1000000000, 0, 1000000000}, // large
		{1, 1, 1000000, 1000000000, 1000000000, 2000000000},
	}
	for _, t := range tests {
		got := minimumCost(t.c1, t.c2, t.cb, t.n1, t.n2)
		status := "OK"
		if got != t.want {
			status = "FAIL"
		}
		fmt.Printf("input=(%d,%d,%d,%d,%d) expected=%d actual=%d %s\n", t.c1, t.c2, t.cb, t.n1, t.n2, t.want, got, status)
	}
}
