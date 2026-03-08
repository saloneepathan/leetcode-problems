package main

import (
	"fmt"
)

func minFlips(s string) int {
	n := len(s)
	s = s + s

	alt1 := make([]byte, 2*n)
	alt2 := make([]byte, 2*n)

	// Create alternating patterns
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			alt1[i] = '0'
			alt2[i] = '1'
		} else {
			alt1[i] = '1'
			alt2[i] = '0'
		}
	}

	res := n
	diff1, diff2 := 0, 0
	left := 0

	for right := 0; right < 2*n; right++ {

		if s[right] != alt1[right] {
			diff1++
		}
		if s[right] != alt2[right] {
			diff2++
		}

		// Maintain window size = n
		if right-left+1 > n {
			if s[left] != alt1[left] {
				diff1--
			}
			if s[left] != alt2[left] {
				diff2--
			}
			left++
		}

		if right-left+1 == n {
			if diff1 < diff2 {
				if diff1 < res {
					res = diff1
				}
			} else {
				if diff2 < res {
					res = diff2
				}
			}
		}
	}

	return res
}

func main() {
	s := "111000"
	fmt.Println("Minimum flips:", minFlips(s))
}
