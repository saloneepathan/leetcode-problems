package main

import "fmt"

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, p := range nums {
		// Special case: 2 is impossible
		if p == 2 {
			ans[i] = -1
			continue
		}

		// Count trailing 1s in p
		t := 0
		for ((p >> t) & 1) == 1 {
			t++
		}

		// Construct the minimum x
		// Keep higher bits same, reduce trailing ones by one
		higher := p >> t
		ans[i] = (higher << t) | ((1 << (t - 1)) - 1)
	}

	return ans
}

func main() {
	// Example test cases
	nums1 := []int{2, 3, 5, 7}
	fmt.Println(minBitwiseArray(nums1)) // [-1 1 4 3]

	nums2 := []int{11, 13, 31}
	fmt.Println(minBitwiseArray(nums2)) // [9 12 15]
}
