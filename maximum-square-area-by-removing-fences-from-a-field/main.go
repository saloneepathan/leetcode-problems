package main

import (
	"fmt"
	"sort"
)

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	const MOD int64 = 1e9 + 7

	// Add boundary fences
	h := append(hFences, 1, m)
	v := append(vFences, 1, n)

	// Sort the fences
	sort.Ints(h)
	sort.Ints(v)

	// Store all horizontal distances
	hDist := make(map[int]bool)
	for i := 0; i < len(h); i++ {
		for j := i + 1; j < len(h); j++ {
			hDist[h[j]-h[i]] = true
		}
	}

	maxSide := 0

	// Find the largest vertical distance also present in horizontal distances
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v); j++ {
			d := v[j] - v[i]
			if hDist[d] && d > maxSide {
				maxSide = d
			}
		}
	}

	if maxSide == 0 {
		return -1
	}

	return int((int64(maxSide) * int64(maxSide)) % MOD)
}

func main() {
	// Example 1
	fmt.Println(maximizeSquareArea(
		4,
		3,
		[]int{2, 3},
		[]int{2},
	)) // Output: 4

	// Example 2
	fmt.Println(maximizeSquareArea(
		6,
		7,
		[]int{2},
		[]int{4},
	)) // Output: -1
}
