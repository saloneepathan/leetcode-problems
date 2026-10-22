package main

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)
	inc, dec := true, true

	for i := 1; i < n; i++ {
		if nums[i] != (nums[i-1]+1)%n {
			inc = false
		}
		if nums[i] != (nums[i-1]-1+n)%n {
			dec = false
		}
	}

	if !inc && !dec {
		return -1
	}

	if inc {
		k := nums[0]
		if k == 0 {
			return 0
		}
		a, b := n-k, k+2
		if a < b {
			return a
		}
		return b
	}

	k := (n - 1 - nums[0] + n) % n
	a, b := k+1, n-k+1
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minOperations([]int{0, 2, 1}))    // 2
	fmt.Println(minOperations([]int{1, 0, 2}))    // 2
	fmt.Println(minOperations([]int{2, 0, 1, 3})) // -1
}
