package main

import (
	"fmt"
)

func pyramidTransition(bottom string, allowed []string) bool {
	// Map from bottom pair -> possible top characters
	adj := make(map[string][]byte)
	for _, a := range allowed {
		key := a[:2]
		adj[key] = append(adj[key], a[2])
	}

	// DFS function
	var dfs func(string) bool
	dfs = func(row string) bool {
		// If we reached the top
		if len(row) == 1 {
			return true
		}

		// Build next level
		var buildNext func(int, []byte) bool
		buildNext = func(index int, next []byte) bool {
			if index == len(row)-1 {
				return dfs(string(next))
			}

			key := row[index : index+2]
			chars, exists := adj[key]
			if !exists {
				return false
			}

			for _, c := range chars {
				next = append(next, c)
				if buildNext(index+1, next) {
					return true
				}
				next = next[:len(next)-1] // backtrack
			}
			return false
		}

		return buildNext(0, []byte{})
	}

	return dfs(bottom)
}

func main() {
	// Example 1
	bottom1 := "BCD"
	allowed1 := []string{"BCC", "CDE", "CEA", "FFF"}
	fmt.Println(pyramidTransition(bottom1, allowed1)) // true

	// Example 2
	bottom2 := "AAAA"
	allowed2 := []string{"AAB", "AAC", "BCD", "BBE", "DEF"}
	fmt.Println(pyramidTransition(bottom2, allowed2)) // false
}
