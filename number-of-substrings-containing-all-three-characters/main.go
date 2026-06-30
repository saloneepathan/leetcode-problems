package main

import "fmt"

func numberOfSubstrings(s string) int {
	n := len(s)
	count := make([]int, 3) // count[0] -> a, count[1] -> b, count[2] -> c

	left := 0
	ans := 0

	for right := 0; right < n; right++ {
		count[s[right]-'a']++

		// Shrink the window while it contains at least one a, b, and c
		for count[0] > 0 && count[1] > 0 && count[2] > 0 {
			ans += n - right
			count[s[left]-'a']--
			left++
		}
	}

	return ans
}

func main() {
	fmt.Println(numberOfSubstrings("abcabc")) // 10
	fmt.Println(numberOfSubstrings("aaacb"))  // 3
	fmt.Println(numberOfSubstrings("abc"))    // 1
}
