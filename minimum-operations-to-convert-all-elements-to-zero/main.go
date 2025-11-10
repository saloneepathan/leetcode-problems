package main

import "fmt"

// minOperations uses a monotonic stack to count the minimal operations
func minOperations(nums []int) int {
	s := []int{}
	res := 0

	for _, a := range nums {
		// Pop higher elements — they can be "covered" by smaller values
		for len(s) > 0 && s[len(s)-1] > a {
			s = s[:len(s)-1]
		}

		// Skip zeros, as they reset the segment
		if a == 0 {
			s = []int{} // clear stack on zero
			continue
		}

		// If stack is empty or top < current, it’s a new unique non-zero layer
		// (means a new operation is needed)
		if len(s) == 0 || s[len(s)-1] < a {
			res++
			s = append(s, a)
		}
	}

	return res
}

func main() {
	fmt.Println(minOperations([]int{0, 2}))                // 1
	fmt.Println(minOperations([]int{3, 1, 2, 1}))          // 3
	fmt.Println(minOperations([]int{1, 2, 1, 2, 1, 2}))    // 4
	fmt.Println(minOperations([]int{4, 4}))                // 1
	fmt.Println(minOperations([]int{7, 2, 0, 4, 2}))       // 4
	fmt.Println(minOperations([]int{0, 0, 0}))             // 0
	fmt.Println(minOperations([]int{5, 1, 1, 2, 0, 3, 3})) // 4
}
