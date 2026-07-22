package main

import (
	"fmt"
	"sort"
)

type SegmentTree struct {
	n   int
	arr []int
	seg []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	if len(arr) == 0 {
		return &SegmentTree{}
	}

	st := &SegmentTree{
		arr: arr,
		n:   len(arr),
		seg: make([]int, len(arr)*4),
	}
	st.build(1, 0, st.n-1)

	return st
}

func (st *SegmentTree) build(p, l, r int) {
	if l == r {
		st.seg[p] = st.arr[l]
		return
	}

	mid := (l + r) >> 1
	st.build(p<<1, l, mid)
	st.build(p<<1|1, mid+1, r)

	st.seg[p] = max(
		st.seg[p<<1],
		st.seg[p<<1|1],
	)
}

func (st *SegmentTree) queryInternal(p, l, r, L, R int) int {
	if L <= l && r <= R {
		return st.seg[p]
	}

	mid := (l + r) >> 1
	res := 0

	if L <= mid {
		res = max(res, st.queryInternal(p<<1, l, mid, L, R))
	}

	if R > mid {
		res = max(res, st.queryInternal(p<<1|1, mid+1, r, L, R))
	}

	return res
}

func (st *SegmentTree) Query(L, R int) int {
	if st.n == 0 || L > R {
		return 0
	}
	return st.queryInternal(1, 0, st.n-1, L, R)
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

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)

	cnt1 := 0
	for _, c := range s {
		if c == '1' {
			cnt1++
		}
	}

	var zeroBlocks []int
	var blockLeft []int
	var blockRight []int

	i := 0
	for i < n {
		start := i

		for i < n && s[i] == s[start] {
			i++
		}

		if s[start] == '0' {
			zeroBlocks = append(zeroBlocks, i-start)
			blockLeft = append(blockLeft, start)
			blockRight = append(blockRight, i-1)
		}
	}

	m := len(zeroBlocks)

	// Less than two zero blocks means no profitable trade.
	if m < 2 {
		ans := make([]int, len(queries))
		for i := range ans {
			ans[i] = cnt1
		}
		return ans
	}

	tmpSum := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		tmpSum[i] = zeroBlocks[i] + zeroBlocks[i+1]
	}

	seg := NewSegmentTree(tmpSum)

	ans := make([]int, len(queries))

	for qIdx, q := range queries {
		l, r := q[0], q[1]

		i := sort.Search(len(blockRight), func(idx int) bool {
			return blockRight[idx] >= l
		})

		j := sort.Search(len(blockLeft), func(idx int) bool {
			return blockLeft[idx] > r
		}) - 1

		// At most one zero block inside the substring.
		if i > m-1 || j < 0 || i >= j {
			ans[qIdx] = cnt1
			continue
		}

		firstLen := blockRight[i] - max(blockLeft[i], l) + 1
		lastLen := min(blockRight[j], r) - blockLeft[j] + 1

		// Exactly two zero blocks.
		if i+1 == j {
			bestGain := firstLen + lastLen
			ans[qIdx] = cnt1 + bestGain
			continue
		}

		val1 := firstLen + zeroBlocks[i+1]
		val2 := zeroBlocks[j-1] + lastLen
		val3 := seg.Query(i+1, j-2)

		bestGain := max(val1, max(val2, val3))

		ans[qIdx] = cnt1 + bestGain
	}

	return ans
}

func main() {
	s := "0100"
	queries := [][]int{
		{0, 3},
		{0, 2},
		{1, 3},
		{2, 3},
	}

	ans := maxActiveSectionsAfterTrade(s, queries)
	fmt.Println(ans)
}
