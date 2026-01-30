package main

import (
	"fmt"
)

func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	const INF int64 = 1e18
	n := len(source)

	// Group rules by substring length
	type Rule struct {
		from string
		to   string
		cost int64
	}
	rulesByLen := map[int][]Rule{}
	for i := range original {
		l := len(original[i])
		rulesByLen[l] = append(rulesByLen[l], Rule{
			from: original[i],
			to:   changed[i],
			cost: int64(cost[i]),
		})
	}

	// Precompute conversion cost maps per length using Floyd–Warshall
	costMap := map[int]map[string]map[string]int64{}

	for l, rules := range rulesByLen {
		nodes := map[string]bool{}
		for _, r := range rules {
			nodes[r.from] = true
			nodes[r.to] = true
		}

		// Initialize distance matrix
		dist := map[string]map[string]int64{}
		for u := range nodes {
			dist[u] = map[string]int64{}
			for v := range nodes {
				if u == v {
					dist[u][v] = 0
				} else {
					dist[u][v] = INF
				}
			}
		}

		// Direct edges
		for _, r := range rules {
			if dist[r.from][r.to] > r.cost {
				dist[r.from][r.to] = r.cost
			}
		}

		// Floyd–Warshall
		for k := range nodes {
			for i := range nodes {
				for j := range nodes {
					if dist[i][k]+dist[k][j] < dist[i][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}

		costMap[l] = dist
	}

	// DP
	dp := make([]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = INF
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		// No operation needed if characters already match
		if source[i] == target[i] {
			dp[i] = dp[i+1]
		}

		// Try all substring operations
		for l, dist := range costMap {
			if i+l > n {
				continue
			}
			sSub := source[i : i+l]
			tSub := target[i : i+l]

			if d1, ok := dist[sSub]; ok {
				if c, ok2 := d1[tSub]; ok2 && c < INF {
					if c+dp[i+l] < dp[i] {
						dp[i] = c + dp[i+l]
					}
				}
			}
		}
	}

	if dp[0] >= INF {
		return -1
	}
	return dp[0]
}

func main() {
	// Example 1
	source1 := "abcd"
	target1 := "acbe"
	original1 := []string{"a", "b", "c", "c", "e", "d"}
	changed1 := []string{"b", "c", "b", "e", "b", "e"}
	cost1 := []int{2, 5, 5, 1, 2, 20}
	fmt.Println(minimumCost(source1, target1, original1, changed1, cost1)) // 28

	// Example 2
	source2 := "abcdefgh"
	target2 := "acdeeghh"
	original2 := []string{"bcd", "fgh", "thh"}
	changed2 := []string{"cde", "thh", "ghh"}
	cost2 := []int{1, 3, 5}
	fmt.Println(minimumCost(source2, target2, original2, changed2, cost2)) // 9

	// Example 3
	source3 := "abcdefgh"
	target3 := "addddddd"
	original3 := []string{"bcd", "defgh"}
	changed3 := []string{"ddd", "ddddd"}
	cost3 := []int{100, 1578}
	fmt.Println(minimumCost(source3, target3, original3, changed3, cost3)) // -1
}
