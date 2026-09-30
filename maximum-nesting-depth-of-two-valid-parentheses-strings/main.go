package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	answer := make([]int, len(seq))
	depth := 0

	for i, ch := range seq {
		if ch == '(' {
			depth++
			answer[i] = depth % 2
		} else {
			answer[i] = depth % 2
			depth--
		}
	}

	return answer
}

func main() {
	seq := "(()())"
	fmt.Println(maxDepthAfterSplit(seq))
}
