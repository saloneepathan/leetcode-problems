package main

import (
	"fmt"
	"sort"
)

// function to count number of 1 bits
func countBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func sortByBits(arr []int) []int {

	sort.Slice(arr, func(i, j int) bool {

		bitsI := countBits(arr[i])
		bitsJ := countBits(arr[j])

		// first sort by number of bits
		if bitsI == bitsJ {
			return arr[i] < arr[j]
		}

		return bitsI < bitsJ
	})

	return arr
}

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	result := sortByBits(arr)

	fmt.Println(result)
}
