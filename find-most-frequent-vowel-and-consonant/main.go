package main

import (
	"fmt"
)

func maxFreqSum(s string) int {
	// arrays for frequency of each letter
	var freq [26]int

	// mark vowels
	vowels := [26]bool{}
	for _, v := range "aeiou" {
		vowels[v-'a'] = true
	}

	maxVowel, maxConsonant := 0, 0

	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		freq[c]++
		if vowels[c] {
			if freq[c] > maxVowel {
				maxVowel = freq[c]
			}
		} else {
			if freq[c] > maxConsonant {
				maxConsonant = freq[c]
			}
		}
	}

	return maxVowel + maxConsonant
}

func main() {
	// Test cases
	fmt.Println(maxFreqSum("successes")) // Expected 6
	fmt.Println(maxFreqSum("aeiaeia"))   // Expected 3
	fmt.Println(maxFreqSum("banana"))    // Vowel 'a'=3, consonant 'n'=2 => 5
	fmt.Println(maxFreqSum("zzz"))       // No vowels, consonant 'z'=3 => 3
	fmt.Println(maxFreqSum("aeiou"))     // Vowels max=1, no consonants => 1
}
