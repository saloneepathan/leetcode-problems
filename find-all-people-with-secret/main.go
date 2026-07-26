package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Pair represents (time, person)
type Pair struct {
	time   int
	person int
}

// MinHeap implementation
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// adjacency list
	adj := make([][]Pair, n)

	for _, meet := range meetings {
		x, y, t := meet[0], meet[1], meet[2]
		adj[x] = append(adj[x], Pair{t, y})
		adj[y] = append(adj[y], Pair{t, x})
	}

	// known[i] = earliest time person i knows the secret
	known := make([]int, n)
	for i := 0; i < n; i++ {
		known[i] = math.MaxInt32
	}

	result := []int{}

	pq := &MinHeap{}
	heap.Init(pq)

	// initial secret holders
	heap.Push(pq, Pair{0, 0})
	heap.Push(pq, Pair{0, firstPerson})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Pair)
		time, person := cur.time, cur.person

		if known[person] != math.MaxInt32 {
			continue
		}

		known[person] = time
		result = append(result, person)

		for _, nei := range adj[person] {
			if known[nei.person] != math.MaxInt32 {
				continue
			}
			if nei.time < time {
				continue
			}
			heap.Push(pq, Pair{nei.time, nei.person})
		}
	}

	return result
}

func main() {
	// Example 1
	n1 := 6
	meetings1 := [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}
	firstPerson1 := 1
	fmt.Println("Output 1:", findAllPeople(n1, meetings1, firstPerson1))

	// Example 2
	n2 := 4
	meetings2 := [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}
	firstPerson2 := 3
	fmt.Println("Output 2:", findAllPeople(n2, meetings2, firstPerson2))

	// Example 3
	n3 := 5
	meetings3 := [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}
	firstPerson3 := 1
	fmt.Println("Output 3:", findAllPeople(n3, meetings3, firstPerson3))
}
