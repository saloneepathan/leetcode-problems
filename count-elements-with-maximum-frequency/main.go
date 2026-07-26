package main

import (
	"fmt"
)

func maxFrequencyElements(nums []int) int {
	// Step 1: Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Step 2: Find max frequency
	maxFreq := 0
	for _, f := range freq {
		if f > maxFreq {
			maxFreq = f
		}
	}

	// Step 3: Sum all elements that appear with max frequency
	result := 0
	for _, f := range freq {
		if f == maxFreq {
			result += f
		}
	}

	return result
}

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 2, 3, 1, 4})) // Output: 4
	fmt.Println(maxFrequencyElements([]int{1, 2, 3, 4, 5}))    // Output: 5
}
