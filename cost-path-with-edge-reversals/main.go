package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a directed edge
type Edge struct {
	to   int
	cost int
}

// State represents a node in the priority queue
type State struct {
	node int
	dist int
}

// MinHeap for Dijkstra
type MinHeap []State

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// minCost finds the minimum cost to travel from node 0 to node n-1
func minCost(n int, edges [][]int) int {
	out := make([][]Edge, n) // outgoing edges
	in := make([][]Edge, n)  // incoming edges

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		out[u] = append(out[u], Edge{v, w})
		in[v] = append(in[v], Edge{u, w})
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt64
	}
	dist[0] = 0

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, State{node: 0, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(State)
		u := cur.node
		d := cur.dist

		if d > dist[u] {
			continue
		}

		// Normal outgoing edges
		for _, e := range out[u] {
			nd := d + e.cost
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, State{node: e.to, dist: nd})
			}
		}

		// Reverse incoming edges using switch at u
		for _, e := range in[u] {
			nd := d + 2*e.cost
			if nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(pq, State{node: e.to, dist: nd})
			}
		}
	}

	if dist[n-1] == math.MaxInt64 {
		return -1
	}
	return dist[n-1]
}

func main() {
	// Example 1
	n1 := 4
	edges1 := [][]int{
		{0, 1, 3},
		{3, 1, 1},
		{2, 3, 4},
		{0, 2, 2},
	}
	fmt.Println(minCost(n1, edges1)) // Output: 5

	// Example 2
	n2 := 4
	edges2 := [][]int{
		{0, 2, 1},
		{2, 1, 1},
		{1, 3, 1},
		{2, 3, 3},
	}
	fmt.Println(minCost(n2, edges2)) // Output: 3
}
