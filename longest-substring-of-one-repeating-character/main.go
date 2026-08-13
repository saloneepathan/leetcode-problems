package main

import "fmt"

type Node struct {
	leftChar  byte
	rightChar byte
	prefix    int
	suffix    int
	best      int
	length    int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(a, b Node) Node {
	if a.length == 0 {
		return b
	}
	if b.length == 0 {
		return a
	}

	res := Node{
		leftChar:  a.leftChar,
		rightChar: b.rightChar,
		prefix:    a.prefix,
		suffix:    b.suffix,
		best:      max(a.best, b.best),
		length:    a.length + b.length,
	}

	// The entire left segment and the beginning of the right
	// segment form one repeating-character prefix.
	if a.leftChar == b.leftChar && a.prefix == a.length {
		res.prefix = a.length + b.prefix
	}

	// The entire right segment and the end of the left
	// segment form one repeating-character suffix.
	if a.rightChar == b.rightChar && b.suffix == b.length {
		res.suffix = b.length + a.suffix
	}

	// A repeating substring may cross the boundary.
	if a.rightChar == b.leftChar {
		res.best = max(res.best, a.suffix+b.prefix)
	}

	return res
}

func build(tree []Node, s []byte, node, left, right int) {
	if left == right {
		tree[node] = Node{
			leftChar:  s[left],
			rightChar: s[left],
			prefix:    1,
			suffix:    1,
			best:      1,
			length:    1,
		}
		return
	}

	mid := (left + right) / 2

	build(tree, s, node*2, left, mid)
	build(tree, s, node*2+1, mid+1, right)

	tree[node] = merge(tree[node*2], tree[node*2+1])
}

func update(tree []Node, node, left, right, idx int, ch byte) {
	if left == right {
		tree[node] = Node{
			leftChar:  ch,
			rightChar: ch,
			prefix:    1,
			suffix:    1,
			best:      1,
			length:    1,
		}
		return
	}

	mid := (left + right) / 2

	if idx <= mid {
		update(tree, node*2, left, mid, idx, ch)
	} else {
		update(tree, node*2+1, mid+1, right, idx, ch)
	}

	tree[node] = merge(tree[node*2], tree[node*2+1])
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	n := len(s)

	tree := make([]Node, 4*n)

	build(tree, []byte(s), 1, 0, n-1)

	ans := make([]int, len(queryIndices))

	for i, idx := range queryIndices {
		update(tree, 1, 0, n-1, idx, queryCharacters[i])
		ans[i] = tree[1].best
	}

	return ans
}

func main() {
	// Example 1
	s := "babacc"
	queryCharacters := "bcb"
	queryIndices := []int{1, 3, 3}

	result := longestRepeating(s, queryCharacters, queryIndices)

	fmt.Println(result) // [3 3 4]

	// Example 2
	s = "abyzz"
	queryCharacters = "aa"
	queryIndices = []int{2, 1}

	result = longestRepeating(s, queryCharacters, queryIndices)

	fmt.Println(result) // [2 3]
}