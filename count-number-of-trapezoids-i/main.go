package main

import (
	"fmt"
)

const MOD = 1_000_000_007

// countTrapezoids returns the number of horizontal trapezoids modulo 1e9+7
func countTrapezoids(points [][]int) int {
	// Count how many points exist at each y-level
	yCount := make(map[int]int)
	for _, p := range points {
		yCount[p[1]]++
	}

	// For each y having k points, compute C(k,2)
	segCounts := make([]int64, 0, len(yCount))
	for _, k := range yCount {
		if k >= 2 {
			v := int64(k)
			segCounts = append(segCounts, (v*(v-1)/2)%MOD)
		}
	}

	if len(segCounts) < 2 {
		return 0
	}

	// Sum of all horizontal segment counts
	var total int64
	for _, v := range segCounts {
		total = (total + v) % MOD
	}

	// Combine every pair: sum(v * (total - v))
	var ans int64
	for _, v := range segCounts {
		ans = (ans + v*(total-v)%MOD) % MOD
	}

	// Each pair counted twice → multiply by modular inverse of 2
	ans = ans * ((MOD + 1) / 2) % MOD

	return int(ans)
}

func main() {
	// Example usage

	points1 := [][]int{
		{1, 0}, {2, 0}, {3, 0},
		{2, 2}, {3, 2},
	}

	points2 := [][]int{
		{0, 0}, {1, 0},
		{0, 1}, {2, 1},
	}

	fmt.Println(countTrapezoids(points1)) // Expected: 3
	fmt.Println(countTrapezoids(points2)) // Expected: 1
}
