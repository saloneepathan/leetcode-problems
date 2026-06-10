package main

import (
	"container/heap"
	"fmt"
	"math/bits"
)

type SparseTable struct {
	lg []int
	mx [][]int
	mn [][]int
}

func NewSparseTable(nums []int) *SparseTable {
	n := len(nums)
	m := bits.Len(uint(n))

	lg := make([]int, n+1)
	for i := 2; i <= n; i++ {
		lg[i] = lg[i/2] + 1
	}

	mx := make([][]int, n)
	mn := make([][]int, n)

	for i := 0; i < n; i++ {
		mx[i] = make([]int, m)
		mn[i] = make([]int, m)
		mx[i][0] = nums[i]
		mn[i][0] = nums[i]
	}

	for j := 1; j < m; j++ {
		for i := 0; i+(1<<j) <= n; i++ {
			mx[i][j] = max(mx[i][j-1], mx[i+(1<<(j-1))][j-1])
			mn[i][j] = min(mn[i][j-1], mn[i+(1<<(j-1))][j-1])
		}
	}

	return &SparseTable{
		lg: lg,
		mx: mx,
		mn: mn,
	}
}

func (st *SparseTable) queryMax(l, r int) int {
	k := st.lg[r-l+1]
	return max(st.mx[l][k], st.mx[r-(1<<k)+1][k])
}

func (st *SparseTable) queryMin(l, r int) int {
	k := st.lg[r-l+1]
	return min(st.mn[l][k], st.mn[r-(1<<k)+1][k])
}

func (st *SparseTable) value(l, r int) int64 {
	return int64(st.queryMax(l, r) - st.queryMin(l, r))
}

type Node struct {
	val int64
	l   int
	r   int
}

type MaxHeap []*Node

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)

	st := NewSparseTable(nums)

	h := &MaxHeap{}
	heap.Init(h)

	for l := 0; l < n; l++ {
		heap.Push(h, &Node{
			val: st.value(l, n-1),
			l:   l,
			r:   n - 1,
		})
	}

	var ans int64

	for k > 0 {
		cur := heap.Pop(h).(*Node)
		ans += cur.val
		k--

		if cur.r > cur.l {
			nr := cur.r - 1
			heap.Push(h, &Node{
				val: st.value(cur.l, nr),
				l:   cur.l,
				r:   nr,
			})
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxTotalValue([]int{1, 3, 2}, 2))    // 4
	fmt.Println(maxTotalValue([]int{4, 2, 5, 1}, 3)) // 12
}
