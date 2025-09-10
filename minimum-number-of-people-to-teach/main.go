package main

import "fmt"

// minimumTeachings returns the minimum number of users we need to teach
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)

	// Build a map for each user's known languages for quick lookup
	knows := make([]map[int]bool, m)
	for i := 0; i < m; i++ {
		knows[i] = make(map[int]bool)
		for _, lang := range languages[i] {
			knows[i][lang] = true
		}
	}

	// Find all friendships where communication is impossible
	toFix := make(map[int]bool) // set of users needing a new language
	for _, f := range friendships {
		u, v := f[0]-1, f[1]-1 // convert to 0-based
		if !canCommunicate(knows[u], knows[v]) {
			toFix[u] = true
			toFix[v] = true
		}
	}

	// If all friendships are fine, no one needs teaching
	if len(toFix) == 0 {
		return 0
	}

	// Try teaching each language, count how many users in "toFix" don't know it
	ans := m // upper bound
	for lang := 1; lang <= n; lang++ {
		count := 0
		for u := range toFix {
			if !knows[u][lang] {
				count++
			}
		}
		if count < ans {
			ans = count
		}
	}

	return ans
}

// Helper: check if two users share a common language
func canCommunicate(a, b map[int]bool) bool {
	for lang := range a {
		if b[lang] {
			return true
		}
	}
	return false
}

func main() {
	// Example 1
	n1 := 2
	languages1 := [][]int{{1}, {2}, {1, 2}}
	friendships1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(minimumTeachings(n1, languages1, friendships1)) // Output: 1

	// Example 2
	n2 := 3
	languages2 := [][]int{{2}, {1, 3}, {1, 2}, {3}}
	friendships2 := [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}}
	fmt.Println(minimumTeachings(n2, languages2, friendships2)) // Output: 2
}
