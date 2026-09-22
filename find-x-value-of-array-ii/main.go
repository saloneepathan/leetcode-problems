package main

import "fmt"

type Node struct {
	prod int
	cnt  [5]int
}

func merge(a, b Node, k int) Node {
	var res Node

	// Product of the entire merged segment.
	res.prod = (a.prod * b.prod) % k

	// Prefixes completely inside the left segment.
	for r := 0; r < k; r++ {
		res.cnt[r] = a.cnt[r]
	}

	// Prefixes that contain the entire left segment
	// and a non-empty prefix of the right segment.
	for r := 0; r < k; r++ {
		newRemainder := (a.prod * r) % k
		res.cnt[newRemainder] += b.cnt[r]
	}

	return res
}

func build(node, l, r int, nums []int, k int, tree []Node) {
	if l == r {
		v := nums[l] % k

		tree[node].prod = v
		tree[node].cnt[v] = 1
		return
	}

	mid := (l + r) >> 1

	build(node<<1, l, mid, nums, k, tree)
	build(node<<1|1, mid+1, r, nums, k, tree)

	tree[node] = merge(
		tree[node<<1],
		tree[node<<1|1],
		k,
	)
}

func update(node, l, r, pos, value, k int, tree []Node) {
	if l == r {
		tree[node] = Node{}

		v := value % k
		tree[node].prod = v
		tree[node].cnt[v] = 1
		return
	}

	mid := (l + r) >> 1

	if pos <= mid {
		update(node<<1, l, mid, pos, value, k, tree)
	} else {
		update(node<<1|1, mid+1, r, pos, value, k, tree)
	}

	tree[node] = merge(
		tree[node<<1],
		tree[node<<1|1],
		k,
	)
}

func query(node, l, r, ql, qr, k int, tree []Node) Node {
	// Completely inside range.
	if ql <= l && r <= qr {
		return tree[node]
	}

	mid := (l + r) >> 1

	// Completely inside left child.
	if qr <= mid {
		return query(
			node<<1,
			l,
			mid,
			ql,
			qr,
			k,
			tree,
		)
	}

	// Completely inside right child.
	if ql > mid {
		return query(
			node<<1|1,
			mid+1,
			r,
			ql,
			qr,
			k,
			tree,
		)
	}

	// Overlaps both children.
	left := query(
		node<<1,
		l,
		mid,
		ql,
		qr,
		k,
		tree,
	)

	right := query(
		node<<1|1,
		mid+1,
		r,
		ql,
		qr,
		k,
		tree,
	)

	return merge(left, right, k)
}

func resultArray(nums []int, k int, queries [][]int) []int {
	n := len(nums)

	tree := make([]Node, 4*n+5)

	// Build the segment tree.
	build(1, 0, n-1, nums, k, tree)

	ans := make([]int, len(queries))

	for i, q := range queries {
		index := q[0]
		value := q[1]
		start := q[2]
		x := q[3]

		// Update persists for all future queries.
		nums[index] = value

		update(
			1,
			0,
			n-1,
			index,
			value,
			k,
			tree,
		)

		// After removing nums[0:start], every valid
		// remaining array is a non-empty prefix of
		// nums[start:n].
		res := query(
			1,
			0,
			n-1,
			start,
			n-1,
			k,
			tree,
		)

		ans[i] = res.cnt[x]
	}

	return ans
}

func main() {
	// Example 1
	nums1 := []int{1, 2, 3, 4, 5}
	k1 := 3
	queries1 := [][]int{
		{2, 2, 0, 2},
		{3, 3, 3, 0},
		{0, 1, 0, 1},
	}

	fmt.Println("Example 1:", resultArray(nums1, k1, queries1))
	// [2 2 2]

	// Example 2
	nums2 := []int{1, 2, 4, 8, 16, 32}
	k2 := 4
	queries2 := [][]int{
		{0, 2, 0, 2},
		{0, 2, 0, 1},
	}

	fmt.Println("Example 2:", resultArray(nums2, k2, queries2))
	// [1 0]

	// Example 3
	nums3 := []int{1, 1, 2, 1, 1}
	k3 := 2
	queries3 := [][]int{
		{2, 1, 0, 1},
	}

	fmt.Println("Example 3:", resultArray(nums3, k3, queries3))
	// [5]
}
