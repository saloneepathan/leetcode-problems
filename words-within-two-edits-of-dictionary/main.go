package main

import "fmt"

func twoEditWords(queries []string, dictionary []string) []string {
	result := []string{}

	for _, q := range queries {
		for _, d := range dictionary {
			if isWithinTwoEdits(q, d) {
				result = append(result, q)
				break
			}
		}
	}

	return result
}

func isWithinTwoEdits(a, b string) bool {
	diff := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}

	return true
}

func main() {
	queries := []string{"word", "note", "ants", "wood"}
	dictionary := []string{"wood", "joke", "moat"}

	result := twoEditWords(queries, dictionary)
	fmt.Println(result) // Output: ["word","note","wood"]
}
