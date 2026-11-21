package main

import (
	"fmt"
)

func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)

	// Build original tree
	children := make([][]int, n)

	for i := 1; i < n; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	// newParent stores the parent after all simultaneous changes.
	newParent := make([]int, n)
	copy(newParent, parent)

	// closest[c] = closest ancestor having character c
	closest := make([]int, 26)

	// -1 means no such ancestor exists.
	for i := range closest {
		closest[i] = -1
	}

	// DFS on the ORIGINAL tree.
	var findNewParents func(int)

	findNewParents = func(node int) {
		c := int(s[node] - 'a')

		// Save the previous closest ancestor.
		oldClosest := closest[c]

		// If an ancestor with the same character exists,
		// it becomes the new parent.
		if oldClosest != -1 {
			newParent[node] = oldClosest
		}

		// Current node becomes the closest ancestor for this character.
		closest[c] = node

		// Continue through the original tree.
		for _, child := range children[node] {
			findNewParents(child)
		}

		// Restore when going back up the DFS.
		closest[c] = oldClosest
	}

	findNewParents(0)

	// Build the FINAL tree using newParent.
	finalChildren := make([][]int, n)

	for i := 1; i < n; i++ {
		finalChildren[newParent[i]] = append(finalChildren[newParent[i]], i)
	}

	// Calculate subtree sizes in the final tree.
	answer := make([]int, n)

	var calculateSize func(int) int

	calculateSize = func(node int) int {
		size := 1

		for _, child := range finalChildren[node] {
			size += calculateSize(child)
		}

		answer[node] = size
		return size
	}

	calculateSize(0)

	return answer
}

func main() {
	parent1 := []int{-1, 0, 0, 1, 1, 1}
	s1 := "abaabc"

	fmt.Println(findSubtreeSizes(parent1, s1))
	// [6 3 1 1 1 1]

	parent2 := []int{-1, 0, 4, 0, 1}
	s2 := "abbba"

	fmt.Println(findSubtreeSizes(parent2, s2))
	// [5 2 1 1 1]
}
