package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
Min-heap of available room numbers
*/
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/*
Min-heap of busy rooms
Sorted by end time, then room number
*/
type BusyRoom struct {
	end  int
	room int
}

type BusyHeap []BusyRoom

func (h BusyHeap) Len() int { return len(h) }
func (h BusyHeap) Less(i, j int) bool {
	if h[i].end == h[j].end {
		return h[i].room < h[j].room
	}
	return h[i].end < h[j].end
}
func (h BusyHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *BusyHeap) Push(x interface{}) { *h = append(*h, x.(BusyRoom)) }
func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/*
Main logic
*/
func mostBooked(n int, meetings [][]int) int {
	// Sort meetings by start time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	available := &IntHeap{}
	busy := &BusyHeap{}
	heap.Init(available)
	heap.Init(busy)

	// Initially all rooms are free
	for i := 0; i < n; i++ {
		heap.Push(available, i)
	}

	count := make([]int, n)

	for _, m := range meetings {
		start, end := m[0], m[1]
		duration := end - start

		// Free rooms that have finished before current meeting starts
		for busy.Len() > 0 && (*busy)[0].end <= start {
			room := heap.Pop(busy).(BusyRoom).room
			heap.Push(available, room)
		}

		if available.Len() > 0 {
			// Assign immediately
			room := heap.Pop(available).(int)
			heap.Push(busy, BusyRoom{end, room})
			count[room]++
		} else {
			// Delay meeting
			br := heap.Pop(busy).(BusyRoom)
			newEnd := br.end + duration
			heap.Push(busy, BusyRoom{newEnd, br.room})
			count[br.room]++
		}
	}

	// Find room with maximum meetings
	maxMeetings := 0
	result := 0
	for i := 0; i < n; i++ {
		if count[i] > maxMeetings {
			maxMeetings = count[i]
			result = i
		}
	}

	return result
}

/*
Example usage
*/
func main() {
	n1 := 2
	meetings1 := [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}
	fmt.Println(mostBooked(n1, meetings1)) // Output: 0

	n2 := 3
	meetings2 := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	fmt.Println(mostBooked(n2, meetings2)) // Output: 1
}
