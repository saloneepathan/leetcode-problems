package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	m, n := len(robot), len(factory)

	// dp[i][j] = min cost to fix robots[i:] using factories[j:]
	dp := make([][]int64, m+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64 / 2
		}
	}

	// base case: no robots left → 0 cost
	for j := 0; j <= n; j++ {
		dp[m][j] = 0
	}

	// fill DP bottom-up
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			// option 1: skip this factory
			dp[i][j] = dp[i][j+1]

			pos := factory[j][0]
			limit := factory[j][1]

			var cost int64 = 0

			// try assigning k robots
			for k := 1; k <= limit && i+k-1 < m; k++ {
				cost += int64(abs(robot[i+k-1] - pos))
				dp[i][j] = min(dp[i][j], cost+dp[i+k][j+1])
			}
		}
	}

	return dp[0][0]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTotalDistance(
		[]int{726554621, -235727278, -199823369},
		[][]int{
			{612684362, 1},
			{519972143, 1},
			{759430060, 2},
			{-76130291, 1},
			{547454631, 2},
			{47263647, 2},
			{-79806151, 2},
			{-329855292, 0},
			{-954058831, 3},
		},
	)) // 308813784 ✅
}
