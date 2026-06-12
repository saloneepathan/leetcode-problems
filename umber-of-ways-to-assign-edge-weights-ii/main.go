package main

import (
	"fmt"
)

const MOD = 1000000007

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1

	// Build tree
	g := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	// Compute LOG
	LOG := 1
	for (1 << LOG) <= n {
		LOG++
	}

	depth := make([]int, n+1)

	up := make([][]int, LOG)
	for i := range up {
		up[i] = make([]int, n+1)
	}

	// Iterative DFS from root = 1
	stack := []int{1}
	visited := make([]bool, n+1)
	visited[1] = true

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, v := range g[u] {
			if visited[v] {
				continue
			}
			visited[v] = true

			depth[v] = depth[u] + 1
			up[0][v] = u

			stack = append(stack, v)
		}
	}

	// Binary lifting table
	for k := 1; k < LOG; k++ {
		for v := 1; v <= n; v++ {
			up[k][v] = up[k-1][up[k-1][v]]
		}
	}

	lca := func(a, b int) int {
		if depth[a] < depth[b] {
			a, b = b, a
		}

		diff := depth[a] - depth[b]

		for k := 0; k < LOG; k++ {
			if ((diff >> k) & 1) == 1 {
				a = up[k][a]
			}
		}

		if a == b {
			return a
		}

		for k := LOG - 1; k >= 0; k-- {
			if up[k][a] != up[k][b] {
				a = up[k][a]
				b = up[k][b]
			}
		}

		return up[0][a]
	}

	// Precompute powers of 2 modulo MOD
	pow2 := make([]int, n)
	pow2[0] = 1

	for i := 1; i < n; i++ {
		pow2[i] = int((int64(pow2[i-1]) * 2) % MOD)
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		u, v := q[0], q[1]

		p := lca(u, v)
		dist := depth[u] + depth[v] - 2*depth[p]

		if dist == 0 {
			ans[i] = 0
		} else {
			ans[i] = pow2[dist-1]
		}
	}

	return ans
}

func main() {
	edges1 := [][]int{
		{1, 2},
	}

	queries1 := [][]int{
		{1, 1},
		{1, 2},
	}

	fmt.Println(assignEdgeWeights(edges1, queries1))
	// [0 1]

	edges2 := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	}

	queries2 := [][]int{
		{1, 4},
		{3, 4},
		{2, 5},
	}

	fmt.Println(assignEdgeWeights(edges2, queries2))
	// [2 1 4]
}
