package main

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])
	count := 0

	for col := 0; col < m; col++ {
		for row := 0; row < n-1; row++ {
			if strs[row][col] > strs[row+1][col] {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"})) // 1
	fmt.Println(minDeletionSize([]string{"a", "b"}))            // 0
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"})) // 3
}
