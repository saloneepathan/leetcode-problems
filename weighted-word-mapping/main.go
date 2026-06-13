package main

import (
	"fmt"
)

func mapWordWeights(words []string, weights []int) string {
	ans := make([]byte, len(words))

	for i, word := range words {
		sum := 0

		for _, ch := range word {
			sum += weights[ch-'a']
		}

		mod := sum % 26

		// Reverse alphabetical mapping:
		// 0 -> 'z', 1 -> 'y', ..., 25 -> 'a'
		ans[i] = byte('z' - mod)
	}

	return string(ans)
}

func main() {
	words1 := []string{"abcd", "def", "xyz"}
	weights1 := []int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2}
	fmt.Println(mapWordWeights(words1, weights1)) // rij

	words2 := []string{"a", "b", "c"}
	weights2 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(mapWordWeights(words2, weights2)) // yyy

	words3 := []string{"abcd"}
	weights3 := []int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5}
	fmt.Println(mapWordWeights(words3, weights3)) // g
}
