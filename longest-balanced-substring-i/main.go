package main

import (
	"fmt"
)

func longestBalanced(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		cnt := make([]int, 26)

		for j := i; j < n; j++ {
			c := s[j] - 'a'
			cnt[c]++
			flag := true

			for _, x := range cnt {
				if x > 0 && x != cnt[c] {
					flag = false
					break
				}
			}

			if flag && (j-i+1) > res {
				res = j - i + 1
			}
		}
	}
	return res
}

func main() {
	var s string
	fmt.Print("Enter a lowercase string: ")
	fmt.Scanln(&s)

	result := longestBalanced(s)
	fmt.Println("Length of longest balanced substring:", result)
}
