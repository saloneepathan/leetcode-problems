package main

import "fmt"

func assignEdgeWeights(edges [][]int) int {
	const MOD = 1000000007

	n := len(edges) + 1

	g := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	maxDepth := 0

	var dfs func(int, int, int)
	dfs = func(u, parent, depth int) {
		if depth > maxDepth {
			maxDepth = depth
		}

		for _, v := range g[u] {
			if v != parent {
				dfs(v, u, depth+1)
			}
		}
	}

	dfs(1, 0, 0)

	// Number of odd-sum assignments for a path of length maxDepth:
	// 2^(maxDepth-1)
	return modPow(2, maxDepth-1, MOD)
}

func modPow(base, exp, mod int) int {
	res := 1

	for exp > 0 {
		if exp&1 == 1 {
			res = res * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}

	return res
}

func main() {
	fmt.Println(assignEdgeWeights([][]int{
		{1, 2},
	})) // 1

	fmt.Println(assignEdgeWeights([][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5},
	})) // 2
}
