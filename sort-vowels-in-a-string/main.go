package main

import (
	"fmt"
)

func sortVowels(s string) string {
	isVowel := func(ch byte) bool {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u',
			'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	// count vowels (ASCII limited)
	freq := make([]int, 128)
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			freq[s[i]]++
		}
	}

	// helper: get next available vowel in sorted ASCII order
	nextVowel := func() byte {
		for i := 0; i < 128; i++ {
			if freq[i] > 0 {
				freq[i]--
				return byte(i)
			}
		}
		return 0
	}

	// rebuild
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			result[i] = nextVowel()
		} else {
			result[i] = s[i]
		}
	}

	return string(result)
}

func main() {
	fmt.Println(sortVowels("lEetcOde")) // "lEOtcede"
	fmt.Println(sortVowels("lYmpH"))    // "lYmpH"
}
