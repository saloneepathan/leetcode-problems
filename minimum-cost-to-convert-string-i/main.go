package main

import (
	"fmt"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const INF int64 = 1e18

	// Distance matrix for 26 lowercase letters
	dist := make([][]int64, 26)
	for i := 0; i < 26; i++ {
		dist[i] = make([]int64, 26)
		for j := 0; j < 26; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	// Direct conversion costs
	for i := 0; i < len(original); i++ {
		u := int(original[i] - 'a')
		v := int(changed[i] - 'a')
		c := int64(cost[i])
		if c < dist[u][v] {
			dist[u][v] = c
		}
	}

	// Floyd–Warshall
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Calculate total cost
	var total int64 = 0
	for i := 0; i < len(source); i++ {
		s := int(source[i] - 'a')
		t := int(target[i] - 'a')

		if dist[s][t] == INF {
			return -1
		}
		total += dist[s][t]
	}

	return total
}

func main() {
	// Example 1
	source := "abcd"
	target := "acbe"
	original := []byte{'a', 'b', 'c', 'c', 'e', 'd'}
	changed := []byte{'b', 'c', 'b', 'e', 'b', 'e'}
	cost := []int{2, 5, 5, 1, 2, 20}

	result := minimumCost(source, target, original, changed, cost)
	fmt.Println("Output:", result) // Expected: 28

	// Example 2
	source2 := "aaaa"
	target2 := "bbbb"
	original2 := []byte{'a', 'c'}
	changed2 := []byte{'c', 'b'}
	cost2 := []int{1, 2}

	result2 := minimumCost(source2, target2, original2, changed2, cost2)
	fmt.Println("Output:", result2) // Expected: 12

	// Example 3
	source3 := "abcd"
	target3 := "abce"
	original3 := []byte{'a'}
	changed3 := []byte{'e'}
	cost3 := []int{10000}

	result3 := minimumCost(source3, target3, original3, changed3, cost3)
	fmt.Println("Output:", result3) // Expected: -1
}
