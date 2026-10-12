package main

import "fmt"

func maxSum(nums []int, k int) int {
	const MOD int64 = 1_000_000_007

	var bits [32]int

	// Count set bits at each position.
	for _, num := range nums {
		for b := 0; b < 32; b++ {
			if (num & (1 << b)) != 0 {
				bits[b]++
			}
		}
	}

	var ans int64

	// Greedily construct the maximum possible numbers.
	for i := 0; i < k; i++ {
		var num int64

		for b := 31; b >= 0; b-- {
			if bits[b] > 0 {
				num |= int64(1) << b
				bits[b]--
			}
		}

		ans = (ans + num*num%MOD) % MOD
	}

	return int(ans)
}

func main() {
	nums1 := []int{2, 6, 5, 8}
	k1 := 2
	fmt.Println(maxSum(nums1, k1)) // Output: 261

	nums2 := []int{4, 5, 4, 7}
	k2 := 3
	fmt.Println(maxSum(nums2, k2)) // Output: 90
}
