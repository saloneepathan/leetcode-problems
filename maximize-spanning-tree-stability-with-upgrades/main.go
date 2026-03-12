package main

import (
	"fmt"
	"sort"
)

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &DSU{parent: parent, rank: rank}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(x, y int) bool {
	px := d.find(x)
	py := d.find(y)

	if px == py {
		return false
	}

	if d.rank[px] < d.rank[py] {
		d.parent[px] = py
	} else if d.rank[px] > d.rank[py] {
		d.parent[py] = px
	} else {
		d.parent[py] = px
		d.rank[px]++
	}

	return true
}

func maxStability(n int, edges [][]int, k int) int {

	mustEdges := [][]int{}
	optEdges := [][]int{}

	for _, e := range edges {
		if e[3] == 1 {
			mustEdges = append(mustEdges, e)
		} else {
			optEdges = append(optEdges, e)
		}
	}

	if len(mustEdges) > n-1 {
		return -1
	}

	dsuInit := NewDSU(n)
	selectedInit := 0
	mustMin := int(1e9)

	for _, e := range mustEdges {
		u, v, s := e[0], e[1], e[2]

		if !dsuInit.union(u, v) {
			return -1
		}

		selectedInit++
		if s < mustMin {
			mustMin = s
		}
	}

	sort.Slice(optEdges, func(i, j int) bool {
		return optEdges[i][2] > optEdges[j][2]
	})

	check := func(target int) bool {

		dsu := &DSU{
			parent: append([]int(nil), dsuInit.parent...),
			rank:   append([]int(nil), dsuInit.rank...),
		}

		selected := selectedInit
		usedUpgrades := 0

		for _, e := range optEdges {

			u, v, s := e[0], e[1], e[2]

			if dsu.find(u) == dsu.find(v) {
				continue
			}

			if s >= target {
				dsu.union(u, v)
				selected++

			} else if usedUpgrades < k && s*2 >= target {
				usedUpgrades++
				dsu.union(u, v)
				selected++
			}

			if selected == n-1 {
				return true
			}
		}

		return false
	}

	l := 0
	r := mustMin
	ans := -1

	for l <= r {

		mid := (l + r) / 2

		if check(mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans
}

func main() {

	n1 := 3
	edges1 := [][]int{
		{0, 1, 2, 1},
		{1, 2, 3, 0},
	}
	k1 := 1

	fmt.Println(maxStability(n1, edges1, k1)) // 2

	n2 := 3
	edges2 := [][]int{
		{0, 1, 4, 0},
		{1, 2, 3, 0},
		{0, 2, 1, 0},
	}
	k2 := 2

	fmt.Println(maxStability(n2, edges2, k2)) // 6

	n3 := 3
	edges3 := [][]int{
		{0, 1, 1, 1},
		{1, 2, 1, 1},
		{2, 0, 1, 1},
	}
	k3 := 0

	fmt.Println(maxStability(n3, edges3, k3)) // -1
}
