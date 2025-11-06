package main

import (
	"container/heap"
	"fmt"
)

// ---------- DSU (Disjoint Set Union) ----------
type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n+1)
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
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

// ---------- Min-Heap for each grid ----------
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

// ---------- Main Function ----------
func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c)

	// Step 1: Build DSU from connections
	for _, conn := range connections {
		dsu.Union(conn[0], conn[1])
	}

	// Step 2: Build heaps for each component
	heaps := make(map[int]*MinHeap)
	online := make([]bool, c+1)
	for i := 1; i <= c; i++ {
		online[i] = true
	}

	for i := 1; i <= c; i++ {
		root := dsu.Find(i)
		if _, ok := heaps[root]; !ok {
			heaps[root] = &MinHeap{}
		}
		heap.Push(heaps[root], i)
	}

	// Step 3: Process queries
	res := []int{}

	for _, q := range queries {
		t, x := q[0], q[1]

		if t == 1 { // Maintenance check
			root := dsu.Find(x)
			if online[x] {
				res = append(res, x)
				continue
			}

			h := heaps[root]
			for h.Len() > 0 {
				top := (*h)[0]
				if online[top] {
					res = append(res, top)
					break
				}
				heap.Pop(h) // lazy delete offline ones
			}
			if h.Len() == 0 {
				res = append(res, -1)
			}

		} else if t == 2 { // Station goes offline
			online[x] = false
		}
	}

	return res
}

// ---------- Example Usage ----------
func main() {
	c := 5
	connections := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	queries := [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
	fmt.Println(processQueries(c, connections, queries)) // Output: [3 2 3]
}
