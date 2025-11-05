package main

import (
	"fmt"
	"sort"
)

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	result := make([]int64, 0, n-k+1)
	freq := make(map[int]int)

	// Helper to compute x-sum for current window
	getXSum := func() int64 {
		type pair struct {
			num, count int
		}
		arr := make([]pair, 0, len(freq))
		for num, cnt := range freq {
			arr = append(arr, pair{num, cnt})
		}

		// Sort by frequency DESC, then by num DESC
		sort.Slice(arr, func(i, j int) bool {
			if arr[i].count == arr[j].count {
				return arr[i].num > arr[j].num
			}
			return arr[i].count > arr[j].count
		})

		sum := int64(0)
		for i := 0; i < len(arr) && i < x; i++ {
			sum += int64(arr[i].num * arr[i].count)
		}
		return sum
	}

	// Initialize first window
	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}
	result = append(result, getXSum())

	// Slide the window
	for i := k; i < n; i++ {
		out := nums[i-k]
		freq[out]--
		if freq[out] == 0 {
			delete(freq, out)
		}
		freq[nums[i]]++
		result = append(result, getXSum())
	}

	return result
}

func main() {
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2)) // [6 10 12]
	fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))       // [11 15 15 15 12]
}
