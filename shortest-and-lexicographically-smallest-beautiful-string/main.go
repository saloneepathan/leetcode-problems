package main

import "fmt"

func shortestBeautifulSubstring(s string, k int) string {
	n := len(s)
	best := ""
	minLen := n + 1

	// Store positions of all 1s.
	ones := []int{}

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ones = append(ones, i)
		}
	}

	// Not enough 1s.
	if len(ones) < k {
		return ""
	}

	// Consider every group of k consecutive 1s.
	for i := 0; i+k-1 < len(ones); i++ {
		start := ones[i]
		end := ones[i+k-1]

		length := end - start + 1
		candidate := s[start : end+1]

		if length < minLen {
			minLen = length
			best = candidate
		} else if length == minLen && candidate < best {
			best = candidate
		}
	}

	return best
}

func main() {
	// Example 1
	s1 := "100011001"
	k1 := 3
	fmt.Println(shortestBeautifulSubstring(s1, k1))
	// Output: 11001

	// Example 2
	s2 := "1011"
	k2 := 2
	fmt.Println(shortestBeautifulSubstring(s2, k2))
	// Output: 11

	// Example 3
	s3 := "000"
	k3 := 1
	fmt.Println(shortestBeautifulSubstring(s3, k3))
	// Output: ""
}