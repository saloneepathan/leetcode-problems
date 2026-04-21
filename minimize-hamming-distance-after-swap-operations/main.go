package main

import "fmt"

// Union-Find structure
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return
	}
	if d.rank[px] < d.rank[py] {
		d.parent[px] = py
	} else if d.rank[px] > d.rank[py] {
		d.parent[py] = px
	} else {
		d.parent[py] = px
		d.rank[px]++
	}
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	dsu := NewDSU(n)

	// Step 1: Build connected components
	for _, swap := range allowedSwaps {
		dsu.Union(swap[0], swap[1])
	}

	// Step 2: Group indices by root
	groups := make(map[int][]int)
	for i := 0; i < n; i++ {
		root := dsu.Find(i)
		groups[root] = append(groups[root], i)
	}

	// Step 3: Calculate minimum Hamming distance
	result := 0

	for _, indices := range groups {
		freq := make(map[int]int)

		// Count source values
		for _, idx := range indices {
			freq[source[idx]]++
		}

		// Match target values
		for _, idx := range indices {
			if freq[target[idx]] > 0 {
				freq[target[idx]]--
			} else {
				result++
			}
		}
	}

	return result
}

// Example usage
func main() {
	source := []int{1, 2, 3, 4}
	target := []int{2, 1, 4, 5}
	allowedSwaps := [][]int{{0, 1}, {2, 3}}

	fmt.Println(minimumHammingDistance(source, target, allowedSwaps)) // Output: 1
}