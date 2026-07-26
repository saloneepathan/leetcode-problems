package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	memoryLimit int
	queue       *list.List      // FIFO queue
	set         map[[3]int]bool // to detect duplicates
	destMap     map[int][]int   // destination -> sorted timestamps
}

// Constructor initializes the Router
func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit: memoryLimit,
		queue:       list.New(),
		set:         make(map[[3]int]bool),
		destMap:     make(map[int][]int),
	}
}

// AddPacket adds a packet if not duplicate
func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if this.set[key] {
		return false
	}

	// Add new packet
	packet := Packet{source, destination, timestamp}
	this.queue.PushBack(packet)
	this.set[key] = true
	this.destMap[destination] = append(this.destMap[destination], timestamp)

	// Evict oldest if memory exceeded
	if this.queue.Len() > this.memoryLimit {
		front := this.queue.Front()
		oldPacket := front.Value.(Packet)
		this.queue.Remove(front)
		delete(this.set, [3]int{oldPacket.source, oldPacket.destination, oldPacket.timestamp})

		// Remove timestamp from destMap
		timestamps := this.destMap[oldPacket.destination]
		idx := sort.SearchInts(timestamps, oldPacket.timestamp)
		if idx < len(timestamps) && timestamps[idx] == oldPacket.timestamp {
			this.destMap[oldPacket.destination] = append(timestamps[:idx], timestamps[idx+1:]...)
		}
	}
	return true
}

// ForwardPacket forwards the next packet in FIFO
func (this *Router) ForwardPacket() []int {
	if this.queue.Len() == 0 {
		return []int{}
	}

	front := this.queue.Front()
	packet := front.Value.(Packet)
	this.queue.Remove(front)
	delete(this.set, [3]int{packet.source, packet.destination, packet.timestamp})

	// Remove timestamp from destMap
	timestamps := this.destMap[packet.destination]
	idx := sort.SearchInts(timestamps, packet.timestamp)
	if idx < len(timestamps) && timestamps[idx] == packet.timestamp {
		this.destMap[packet.destination] = append(timestamps[:idx], timestamps[idx+1:]...)
	}

	return []int{packet.source, packet.destination, packet.timestamp}
}

// GetCount returns number of packets with given destination and timestamp in range
func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	timestamps, exists := this.destMap[destination]
	if !exists || len(timestamps) == 0 {
		return 0
	}

	// Find first index >= startTime
	left := sort.Search(len(timestamps), func(i int) bool {
		return timestamps[i] >= startTime
	})
	// Find first index > endTime
	right := sort.Search(len(timestamps), func(i int) bool {
		return timestamps[i] > endTime
	})

	return right - left
}

// ==================== TESTING ====================
func main() {
	// Example 1
	fmt.Println("Example 1:")
	router := Constructor(3)
	fmt.Println(router.AddPacket(1, 4, 90))   // true
	fmt.Println(router.AddPacket(2, 5, 90))   // true
	fmt.Println(router.AddPacket(1, 4, 90))   // false (duplicate)
	fmt.Println(router.AddPacket(3, 5, 95))   // true
	fmt.Println(router.AddPacket(4, 5, 105))  // true (evicts [1,4,90])
	fmt.Println(router.ForwardPacket())       // [2 5 90]
	fmt.Println(router.AddPacket(5, 2, 110))  // true
	fmt.Println(router.GetCount(5, 100, 110)) // 1

	// Example 2
	fmt.Println("\nExample 2:")
	router2 := Constructor(2)
	fmt.Println(router2.AddPacket(7, 4, 90)) // true
	fmt.Println(router2.ForwardPacket())     // [7 4 90]
	fmt.Println(router2.ForwardPacket())     // []
}
