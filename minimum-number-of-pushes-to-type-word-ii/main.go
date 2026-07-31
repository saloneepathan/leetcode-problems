package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	freq := make([]int, 26)

	for _, ch := range word {
		freq[ch-'a']++
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	ans := 0
	for i, f := range freq {
		if f == 0 {
			break
		}

		// First 8 letters -> 1 push
		// Next 8 letters -> 2 pushes
		// Next 8 letters -> 3 pushes
		// Remaining letters -> 4 pushes
		pushes := i/8 + 1
		ans += f * pushes
	}

	return ans
}

func main() {
	fmt.Println(minimumPushes("abcde"))                  // 5
	fmt.Println(minimumPushes("xyzxyzxyzxyz"))           // 12
	fmt.Println(minimumPushes("aabbccddeeffgghhiiiiii")) // 24
}
