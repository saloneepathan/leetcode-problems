package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}

	// Create a sorted copy of the array.
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)

	// Assign ranks to unique values.
	rank := make(map[int]int)
	currentRank := 1
	for _, v := range sorted {
		if _, exists := rank[v]; !exists {
			rank[v] = currentRank
			currentRank++
		}
	}

	// Build the result.
	result := make([]int, n)
	for i, v := range arr {
		result[i] = rank[v]
	}

	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	result := arrayRankTransform(arr)

	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
