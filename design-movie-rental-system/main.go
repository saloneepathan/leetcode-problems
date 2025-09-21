package main

import (
	"container/heap"
	"fmt"
)

// ----------------- Implementation -----------------

type MovieRentingSystem struct {
	price    map[[2]int]int  // (shop,movie) -> price
	state    map[[2]int]int  // version counter for each (shop,movie)
	isRented map[[2]int]bool // current rented state
	avail    map[int]*AvailHeap
	rentedH  *RentedHeap
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	m := MovieRentingSystem{
		price:    make(map[[2]int]int),
		state:    make(map[[2]int]int),
		isRented: make(map[[2]int]bool),
		avail:    make(map[int]*AvailHeap),
		rentedH:  &RentedHeap{},
	}
	heap.Init(m.rentedH)
	for _, e := range entries {
		shop, movie, p := e[0], e[1], e[2]
		key := [2]int{shop, movie}
		m.price[key] = p
		m.state[key] = 1
		if m.avail[movie] == nil {
			m.avail[movie] = &AvailHeap{}
			heap.Init(m.avail[movie])
		}
		heap.Push(m.avail[movie], AvailNode{price: p, shop: shop, ver: 1})
	}
	return m
}

func (this *MovieRentingSystem) Search(movie int) []int {
	h := this.avail[movie]
	if h == nil || h.Len() == 0 {
		return []int{}
	}
	res := []int{}
	temp := []AvailNode{}
	for h.Len() > 0 && len(res) < 5 {
		node := heap.Pop(h).(AvailNode)
		key := [2]int{node.shop, movie}
		if node.ver == this.state[key] && !this.isRented[key] {
			res = append(res, node.shop)
			temp = append(temp, node)
		}
	}
	for _, node := range temp {
		heap.Push(h, node)
	}
	return res
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	key := [2]int{shop, movie}
	this.isRented[key] = true
	this.state[key]++
	ver := this.state[key]
	p := this.price[key]
	heap.Push(this.rentedH, RentedNode{price: p, shop: shop, movie: movie, ver: ver})
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	key := [2]int{shop, movie}
	this.isRented[key] = false
	this.state[key]++
	ver := this.state[key]
	p := this.price[key]
	if this.avail[movie] == nil {
		this.avail[movie] = &AvailHeap{}
		heap.Init(this.avail[movie])
	}
	heap.Push(this.avail[movie], AvailNode{price: p, shop: shop, ver: ver})
}

func (this *MovieRentingSystem) Report() [][]int {
	res := [][]int{}
	temp := []RentedNode{}
	for this.rentedH.Len() > 0 && len(res) < 5 {
		node := heap.Pop(this.rentedH).(RentedNode)
		key := [2]int{node.shop, node.movie}
		if node.ver == this.state[key] && this.isRented[key] {
			res = append(res, []int{node.shop, node.movie})
			temp = append(temp, node)
		}
	}
	for _, node := range temp {
		heap.Push(this.rentedH, node)
	}
	return res
}

// ----------------- Heaps -----------------

type AvailNode struct {
	price int
	shop  int
	ver   int
}
type AvailHeap []AvailNode

func (h AvailHeap) Len() int { return len(h) }
func (h AvailHeap) Less(i, j int) bool {
	if h[i].price == h[j].price {
		return h[i].shop < h[j].shop
	}
	return h[i].price < h[j].price
}
func (h AvailHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *AvailHeap) Push(x interface{}) { *h = append(*h, x.(AvailNode)) }
func (h *AvailHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type RentedNode struct {
	price int
	shop  int
	movie int
	ver   int
}
type RentedHeap []RentedNode

func (h RentedHeap) Len() int { return len(h) }
func (h RentedHeap) Less(i, j int) bool {
	if h[i].price == h[j].price {
		if h[i].shop == h[j].shop {
			return h[i].movie < h[j].movie
		}
		return h[i].shop < h[j].shop
	}
	return h[i].price < h[j].price
}
func (h RentedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *RentedHeap) Push(x interface{}) { *h = append(*h, x.(RentedNode)) }
func (h *RentedHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ----------------- Testing -----------------

func main() {
	entries := [][]int{
		{0, 1, 5}, {0, 2, 6}, {0, 3, 7},
		{1, 1, 4}, {1, 2, 7},
		{2, 1, 5},
	}
	system := Constructor(3, entries)

	fmt.Println(system.Search(1)) // expect [1,0,2]

	system.Rent(0, 1)
	system.Rent(1, 2)

	fmt.Println(system.Report()) // expect [[0,1],[1,2]]

	system.Drop(1, 2)

	fmt.Println(system.Search(2)) // expect [0,1]
}
