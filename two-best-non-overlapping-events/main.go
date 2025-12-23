package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Event represents an event in the heap
type Event struct {
	end   int
	value int
}

// MinHeap ordered by event end time
type MinHeap []Event

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Event)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Solution function
func maxTwoEvents(events [][]int) int {
	// Sort events by start time
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)

	maxPrev := 0
	ans := 0

	for _, e := range events {
		start, end, val := e[0], e[1], e[2]

		// Remove all events that end before current start
		for h.Len() > 0 && (*h)[0].end < start {
			ev := heap.Pop(h).(Event)
			if ev.value > maxPrev {
				maxPrev = ev.value
			}
		}

		// Combine current event with best previous
		if maxPrev+val > ans {
			ans = maxPrev + val
		}

		// Consider taking only this event
		if val > ans {
			ans = val
		}

		// Push current event into heap
		heap.Push(h, Event{end: end, value: val})
	}

	return ans
}

// Main function with test cases
func main() {
	events1 := [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}
	events2 := [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}
	events3 := [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}

	fmt.Println(maxTwoEvents(events1)) // 4
	fmt.Println(maxTwoEvents(events2)) // 5
	fmt.Println(maxTwoEvents(events3)) // 8
}
