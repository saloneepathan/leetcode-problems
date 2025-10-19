package main

import (
	"fmt"
)

// findLexSmallestString returns the lexicographically smallest string
// after applying any number of allowed operations.
func findLexSmallestString(s string, a int, b int) string {
	visited := make(map[string]bool)
	queue := []string{s}
	minStr := s
	n := len(s)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visited[cur] {
			continue
		}
		visited[cur] = true

		// update lexicographically smallest
		if cur < minStr {
			minStr = cur
		}

		// Operation 1: add a to all odd indices
		bytes := []byte(cur)
		for i := 1; i < n; i += 2 {
			bytes[i] = byte((int(bytes[i]-'0')+a)%10 + '0')
		}
		added := string(bytes)
		if !visited[added] {
			queue = append(queue, added)
		}

		// Operation 2: rotate right by b
		rotated := cur[n-b:] + cur[:n-b]
		if !visited[rotated] {
			queue = append(queue, rotated)
		}
	}

	return minStr
}

func main() {
	fmt.Println(findLexSmallestString("5525", 9, 2)) // Output: "2050"
	fmt.Println(findLexSmallestString("74", 5, 1))   // Output: "24"
	fmt.Println(findLexSmallestString("0011", 4, 2)) // Output: "0011"

	// extra tests
	fmt.Println(findLexSmallestString("43987654", 7, 3))
	fmt.Println(findLexSmallestString("123456", 2, 4))
}
