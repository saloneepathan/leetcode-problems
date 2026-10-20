package main

import "fmt"

func sortableIntegers(nums []int) int {
	n := len(nums)

	// Prefix sum of strict descents between adjacent elements.
	prefix := make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1]
		if nums[i-1] > nums[i] {
			prefix[i]++
		}
	}

	isSortable := func(k int) bool {
		for left := 0; left < n; left += k {
			right := left + k - 1

			// Count descents inside this block.
			descents := prefix[right] - prefix[left]

			// Include the circular boundary.
			if nums[right] > nums[left] {
				descents++
			}

			if descents > 1 {
				return false
			}
		}
		return true
	}

	ans := 0

	// Enumerate all divisors of n.
	for k := 1; k*k <= n; k++ {
		if n%k != 0 {
			continue
		}

		if isSortable(k) {
			ans += k
		}

		other := n / k
		if other != k && isSortable(other) {
			ans += other
		}
	}

	return ans
}

func main() {
	fmt.Println(sortableIntegers([]int{3, 1, 2})) // 3
	fmt.Println(sortableIntegers([]int{7, 6, 5})) // 0
	fmt.Println(sortableIntegers([]int{5, 8}))    // 3
}