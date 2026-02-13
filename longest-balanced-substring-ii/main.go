package main

import (
	"fmt"
)

func longestBalanced(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	ans := 1

	type pair struct {
		x, y int
	}

	A, B, C := 0, 0, 0

	abc := make(map[pair]int)
	ab := make(map[pair]int)
	bc := make(map[pair]int)
	ca := make(map[pair]int)

	// Base case: empty prefix
	abc[pair{0, 0}] = -1
	ab[pair{0, 0}] = -1
	bc[pair{0, 0}] = -1
	ca[pair{0, 0}] = -1

	streak := 1

	for i := 0; i < n; i++ {

		// Track single-character streak
		if i > 0 {
			if s[i] == s[i-1] {
				streak++
			} else {
				streak = 1
			}
			if streak > ans {
				ans = streak
			}
		}

		// Update prefix counts
		switch s[i] {
		case 'a':
			A++
		case 'b':
			B++
		case 'c':
			C++
		}

		// Case 1: A = B = C
		key := pair{B - A, C - A}
		if idx, ok := abc[key]; ok {
			if i-idx > ans {
				ans = i - idx
			}
		} else {
			abc[key] = i
		}

		// Case 2: A = B and no C
		key = pair{A - B, C}
		if idx, ok := ab[key]; ok {
			if i-idx > ans {
				ans = i - idx
			}
		} else {
			ab[key] = i
		}

		// Case 3: B = C and no A
		key = pair{B - C, A}
		if idx, ok := bc[key]; ok {
			if i-idx > ans {
				ans = i - idx
			}
		} else {
			bc[key] = i
		}

		// Case 4: C = A and no B
		key = pair{C - A, B}
		if idx, ok := ca[key]; ok {
			if i-idx > ans {
				ans = i - idx
			}
		} else {
			ca[key] = i
		}
	}

	return ans
}

func main() {
	fmt.Println(longestBalanced("a"))     // 1
	fmt.Println(longestBalanced("aa"))    // 2
	fmt.Println(longestBalanced("abbac")) // 4
	fmt.Println(longestBalanced("aabcc")) // 3
	fmt.Println(longestBalanced("aba"))   // 2
}
