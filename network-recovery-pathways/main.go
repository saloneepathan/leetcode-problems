package main

import (
	"fmt"
)

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)

	type Edge struct {
		to   int
		cost int
	}

	graph := make([][]Edge, n)
	indegree := make([]int, n)

	maxCost := 0
	for _, e := range edges {
		u, v, c := e[0], e[1], e[2]
		graph[u] = append(graph[u], Edge{v, c})
		indegree[v]++
		if c > maxCost {
			maxCost = c
		}
	}

	// Topological Sort (Kahn's Algorithm)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := make([]int, 0, n)
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		topo = append(topo, u)

		for _, e := range graph[u] {
			indegree[e.to]--
			if indegree[e.to] == 0 {
				queue = append(queue, e.to)
			}
		}
	}

	const INF int64 = 1 << 60

	check := func(limit int) bool {
		dist := make([]int64, n)
		for i := range dist {
			dist[i] = INF
		}
		dist[0] = 0

		for _, u := range topo {
			if dist[u] == INF {
				continue
			}

			// Intermediate offline nodes cannot be part of the path.
			if u != 0 && u != n-1 && !online[u] {
				continue
			}

			for _, e := range graph[u] {
				v := e.to
				c := e.cost

				if c < limit {
					continue
				}

				// Destination can be offline only if it is n-1.
				if v != n-1 && !online[v] {
					continue
				}

				if dist[u]+int64(c) < dist[v] {
					dist[v] = dist[u] + int64(c)
				}
			}
		}

		return dist[n-1] <= k
	}

	if !check(0) {
		return -1
	}

	lo, hi := 0, maxCost
	ans := 0

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return ans
}

func main() {
	// Example 1
	edges1 := [][]int{
		{0, 1, 5},
		{1, 3, 10},
		{0, 2, 3},
		{2, 3, 4},
	}
	online1 := []bool{true, true, true, true}
	k1 := int64(10)

	fmt.Println(findMaxPathScore(edges1, online1, k1)) // 3

	// Example 2
	edges2 := [][]int{
		{0, 1, 7},
		{1, 4, 5},
		{0, 2, 6},
		{2, 3, 6},
		{3, 4, 2},
		{2, 4, 6},
	}
	online2 := []bool{true, true, true, false, true}
	k2 := int64(12)

	fmt.Println(findMaxPathScore(edges2, online2, k2)) // 6
}
