package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFrequency(nums []int, k int, numOperations int) int {
	// find the maximum element
	maxElem := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxElem {
			maxElem = nums[i]
		}
	}

	// construct a frequency array and a "number line" array
	n := maxElem + k + 2 // ensure index maxElem+k+1 is in bounds
	freqs := make([]int, n)
	numLine := make([]int, n)

	for _, num := range nums {
		freqs[num]++
		left := num - k
		if left < 0 {
			left = 0
		}
		numLine[left]++
		numLine[num+k+1]--
	}

	// compute prefix sums and answer
	ans := 0
	for i := 1; i < n; i++ {
		numLine[i] += numLine[i-1]
		// numLine[i]-freqs[i]-numOperations represents
		// how many nums can't be adjusted to i given allowed operations
		val := numLine[i] - max(0, numLine[i]-freqs[i]-numOperations)
		if val > ans {
			ans = val
		}
	}

	return ans
}

func main() {
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2))       // Expected: 2
	fmt.Println(maxFrequency([]int{5, 11, 20, 20}, 5, 1)) // Expected: 2
	fmt.Println(maxFrequency([]int{1, 90}, 76, 1))        // Expected: 1
	fmt.Println(maxFrequency([]int{2, 2, 2}, 1, 1))       // Expected: 3
}
