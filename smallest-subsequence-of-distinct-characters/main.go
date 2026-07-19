package main

import "fmt"

func smallestSubsequence(s string) string {
	n := len(s)

	last := make([]int, 26)
	for i := 0; i < n; i++ {
		last[s[i]-'a'] = i
	}

	stack := make([]byte, 0)
	visited := make([]bool, 26)

	for i := 0; i < n; i++ {
		ch := s[i]

		if visited[ch-'a'] {
			continue
		}

		for len(stack) > 0 &&
			stack[len(stack)-1] > ch &&
			last[stack[len(stack)-1]-'a'] > i {

			visited[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, ch)
		visited[ch-'a'] = true
	}

	return string(stack)
}

func main() {
	fmt.Println(smallestSubsequence("bcabc"))    // abc
	fmt.Println(smallestSubsequence("cbacdcbc")) // acdb
}
