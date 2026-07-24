package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func uniqueXorTriplets(nums []int) int {
	n := len(nums)

	m := 0
	for _, v := range nums {
		m = max(m, v)
	}

	// Find the smallest power of 2 greater than max element.
	u := 1
	for u <= m {
		u <<= 1
	}

	// Store all possible XORs of two elements (i <= j).
	s := make([]bool, u)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			s[nums[i]^nums[j]] = true
		}
	}

	// XOR each value with every element to form triplets.
	t := make([]bool, u)
	for x := 0; x < u; x++ {
		if !s[x] {
			continue
		}
		for _, v := range nums {
			t[x^v] = true
		}
	}

	// Count unique XOR values.
	ans := 0
	for _, exists := range t {
		if exists {
			ans++
		}
	}

	return ans
}

func main() {
	nums1 := []int{1, 3}
	fmt.Println(uniqueXorTriplets(nums1)) // Output: 2

	nums2 := []int{1, 2, 3}
	fmt.Println(uniqueXorTriplets(nums2))

	nums3 := []int{5, 1, 7, 3}
	fmt.Println(uniqueXorTriplets(nums3))
}
