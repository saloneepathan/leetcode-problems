package main

import (
	"fmt"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

// ===================== MAIN FUNCTION =====================

func findXSum(nums []int, k int, x int) []int64 {
	helper := NewHelper(x)
	ans := []int64{}

	for i := 0; i < len(nums); i++ {
		helper.Insert(nums[i])
		if i >= k {
			helper.Remove(nums[i-k])
		}
		if i >= k-1 {
			ans = append(ans, helper.Get())
		}
	}

	return ans
}

// ===================== HELPER STRUCT =====================

type Helper struct {
	x      int
	result int64
	large  *redblacktree.Tree[pair, struct{}] // top x elements
	small  *redblacktree.Tree[pair, struct{}] // remaining
	occ    map[int]int                        // frequency map
}

type pair struct {
	freq int
	num  int
}

// Custom comparator for ordering pairs:
// higher freq = greater
// if freq equal, higher num = greater
func pairComparator(a, b pair) int {
	if a.freq != b.freq {
		return a.freq - b.freq
	}
	return a.num - b.num
}

func NewHelper(x int) *Helper {
	return &Helper{
		x:      x,
		result: 0,
		large:  redblacktree.NewWith[pair, struct{}](pairComparator),
		small:  redblacktree.NewWith[pair, struct{}](pairComparator),
		occ:    make(map[int]int),
	}
}

// ===================== CORE METHODS =====================

func (h *Helper) Insert(num int) {
	if h.occ[num] > 0 {
		h.internalRemove(pair{freq: h.occ[num], num: num})
	}
	h.occ[num]++
	h.internalInsert(pair{freq: h.occ[num], num: num})
}

func (h *Helper) Remove(num int) {
	if h.occ[num] == 0 {
		return
	}
	h.internalRemove(pair{freq: h.occ[num], num: num})
	h.occ[num]--
	if h.occ[num] > 0 {
		h.internalInsert(pair{freq: h.occ[num], num: num})
	}
}

func (h *Helper) Get() int64 {
	return h.result
}

// ===================== INTERNAL HELPERS =====================

// Insert a pair into the correct tree and update result if in top x
func (h *Helper) internalInsert(p pair) {
	if h.large.Size() < h.x {
		h.result += int64(p.freq) * int64(p.num)
		h.large.Put(p, struct{}{})
	} else {
		minLarge := h.large.Left().Key
		// If new element is more frequent (or larger value on tie)
		if pairComparator(p, minLarge) > 0 {
			h.result += int64(p.freq) * int64(p.num)
			h.large.Put(p, struct{}{})
			toRemove := h.large.Left().Key
			h.result -= int64(toRemove.freq) * int64(toRemove.num)
			h.large.Remove(toRemove)
			h.small.Put(toRemove, struct{}{})
		} else {
			h.small.Put(p, struct{}{})
		}
	}
}

// Remove a pair from one of the trees and rebalance if needed
func (h *Helper) internalRemove(p pair) {
	if _, found := h.large.Get(p); found {
		h.result -= int64(p.freq) * int64(p.num)
		h.large.Remove(p)

		// If large has fewer than x elements, move the best from small
		if h.small.Size() > 0 {
			maxSmall := h.small.Right().Key
			h.result += int64(maxSmall.freq) * int64(maxSmall.num)
			h.small.Remove(maxSmall)
			h.large.Put(maxSmall, struct{}{})
		}
	} else if _, found := h.small.Get(p); found {
		h.small.Remove(p)
	}
}

// ===================== TEST =====================

func main() {
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2)) // [6 10 12]
	fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))       // [11 15 15 15 12]
}
