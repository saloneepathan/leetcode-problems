package main

import (
	"fmt"
	"math"
)

func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	minDist := math.MaxInt32

	for i := 0; i < n; i++ {
		if words[i] == target {
			diff := abs(i - startIndex)
			dist := min(diff, n-diff)
			minDist = min(minDist, dist)
		}
	}

	if minDist == math.MaxInt32 {
		return -1
	}
	return minDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	target := "hello"
	startIndex := 1

	fmt.Println(closetTarget(words, target, startIndex)) // Output: 1
}
