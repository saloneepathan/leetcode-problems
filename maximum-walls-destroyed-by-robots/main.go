package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWalls(robots []int, distance []int, walls []int) int {
	const (
		LL = 0
		LR = 1
		RL = 2
		RR = 3
	)

	// Append sentinel robots
	robots = append(robots, -1)
	robots = append(robots, 2000000000)
	distance = append(distance, 0)
	distance = append(distance, 0)

	n := len(robots)

	// rIndices
	rIndices := make([]int, n)
	for i := 0; i < n; i++ {
		rIndices[i] = i
	}

	sort.Slice(rIndices, func(i, j int) bool {
		return robots[rIndices[i]] < robots[rIndices[j]]
	})

	sort.Ints(walls)

	curr := 0
	lbot := rIndices[curr]
	rbot := rIndices[curr+1]

	dp := make([]int, 4)

	for i := 0; i < len(walls); i++ {
		for walls[i] > robots[rbot] {
			curr++
			lbot = rIndices[curr]
			rbot = rIndices[curr+1]

			maxL := max(dp[LL], dp[RL])
			maxR := max(dp[LR], dp[RR])

			dp[LL] = maxL
			dp[LR] = maxL
			dp[RL] = maxR
			dp[RR] = maxR
		}

		if walls[i] >= robots[rbot]-distance[rbot] {
			dp[LL]++
		}

		if walls[i] == robots[rbot] {
			dp[LR]++
		}

		if walls[i] <= robots[lbot]+distance[lbot] ||
			walls[i] >= robots[rbot]-distance[rbot] {
			dp[RL]++
		}

		if walls[i] <= robots[lbot]+distance[lbot] ||
			walls[i] == robots[rbot] {
			dp[RR]++
		}
	}

	return max(max(dp[0], dp[1]), max(dp[2], dp[3]))
}

func main() {
	robots := []int{1, 4}
	distance := []int{1, 1}
	walls := []int{2, 3, 4}

	fmt.Println(maxWalls(robots, distance, walls))
}
