package main

import "fmt"

// doesAliceWin returns true if Alice can win the game
func doesAliceWin(s string) bool {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			return true // as soon as we see a vowel, Alice can win
		}
	}
	return false // no vowels at all -> Alice loses immediately
}

func main() {
	fmt.Println(doesAliceWin("leetcoder")) // true
	fmt.Println(doesAliceWin("bbcd"))      // false
}
