package main

import "fmt"

func minDeletionSize(strs []string) int {
	n := len(strs)
	if n <= 1 {
		return 0
	}

	m := len(strs[0])
	deletions := 0

	// sorted[i] means strs[i] < strs[i+1] is already confirmed
	sorted := make([]bool, n-1)

	for col := 0; col < m; col++ {
		needDelete := false

		// Check if keeping this column breaks lexicographic order
		for i := 1; i < n; i++ {
			if !sorted[i-1] && strs[i-1][col] > strs[i][col] {
				needDelete = true
				break
			}
		}

		if needDelete {
			deletions++
			continue
		}

		// Update sorted status
		for i := 1; i < n; i++ {
			if !sorted[i-1] && strs[i-1][col] < strs[i][col] {
				sorted[i-1] = true
			}
		}
	}

	return deletions
}

func main() {
	tests := [][]string{
		{"ca", "bb", "ac"},
		{"xc", "yb", "za"},
		{"zyx", "wvu", "tsr"},
		{"abcdef", "uvwxyz"},
	}

	for _, test := range tests {
		fmt.Println("Input:", test)
		fmt.Println("Output:", minDeletionSize(test))
		fmt.Println()
	}
}
