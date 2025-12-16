package main

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	// Build adjacency list
	g := make([][]int, n)
	for _, e := range hierarchy {
		u := e[0] - 1
		v := e[1] - 1
		g[u] = append(g[u], v)
	}

	var dfs func(int) result
	dfs = func(u int) result {
		cost := present[u]
		dCost := present[u] / 2

		// dp0: parent NOT bought
		// dp1: parent bought (discount applies)
		dp0 := make([]int, budget+1)
		dp1 := make([]int, budget+1)

		// subtree profits
		subProfit0 := make([]int, budget+1)
		subProfit1 := make([]int, budget+1)

		uSize := 0

		for _, v := range g[u] {
			child := dfs(v)
			uSize += child.size

			// knapsack merge
			for i := budget; i >= 0; i-- {
				best0 := subProfit0[i]
				best1 := subProfit1[i]
				for sub := 0; sub <= min(child.size, i); sub++ {
					best0 = max(best0, subProfit0[i-sub]+child.dp0[sub])
					best1 = max(best1, subProfit1[i-sub]+child.dp1[sub])
				}
				subProfit0[i] = best0
				subProfit1[i] = best1
			}
		}

		// Decide buying u or not
		for i := 0; i <= budget; i++ {
			dp0[i] = subProfit0[i]
			dp1[i] = subProfit0[i]

			if i >= cost {
				dp0[i] = max(dp0[i], subProfit1[i-cost]+future[u]-cost)
			}
			if i >= dCost {
				dp1[i] = max(dp1[i], subProfit1[i-dCost]+future[u]-dCost)
			}
		}

		return result{
			dp0:  dp0,
			dp1:  dp1,
			size: uSize + cost,
		}
	}

	res := dfs(0)
	return res.dp0[budget]
}

type result struct {
	dp0  []int
	dp1  []int
	size int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 3
	present := []int{5, 2, 3}
	future := []int{8, 5, 6}
	hierarchy := [][]int{{1, 2}, {2, 3}}
	budget := 7

	println(maxProfit(n, present, future, hierarchy, budget)) // Expected: 12
}
