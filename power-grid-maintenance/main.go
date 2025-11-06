package main

import (
	"container/heap"
	"fmt"
)

// ---------- DSU ----------
type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n+1)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	return &DSU{p, r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
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

// ---------- Min-Heap ----------
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- Main ----------
func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c)
	online := make([]bool, c+1)
	for i := 1; i <= c; i++ {
		online[i] = true
	}

	// Build connectivity
	for _, e := range connections {
		dsu.Union(e[0], e[1])
	}

	// Build heaps for each connected component
	heaps := make([]*MinHeap, c+1)
	for i := 1; i <= c; i++ {
		root := dsu.Find(i)
		if heaps[root] == nil {
			heaps[root] = &MinHeap{}
		}
		heap.Push(heaps[root], i)
	}

	res := []int{}

	for _, q := range queries {
		t, x := q[0], q[1]

		if t == 1 { // Maintenance query
			root := dsu.Find(x)
			if online[x] {
				res = append(res, x)
				continue
			}

			h := heaps[root]
			for h != nil && h.Len() > 0 {
				top := (*h)[0]
				if online[top] {
					res = append(res, top)
					break
				}
				heap.Pop(h) // remove offline ones
			}
			if h == nil || h.Len() == 0 {
				res = append(res, -1)
			}

		} else { // t == 2, take station offline
			online[x] = false
		}
	}

	return res
}

// ---------- Example ----------
func main() {
	c := 5
	connections := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	queries := [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
	fmt.Println(processQueries(c, connections, queries)) // [3 2 3]
}
