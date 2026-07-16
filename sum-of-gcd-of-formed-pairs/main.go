package main

import (
	"fmt"
	"sort"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdSum(nums []int) int64 {
	n := len(nums)

	prefixGcd := make([]int, n)

	mx := 0
	for i := 0; i < n; i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
		prefixGcd[i] = gcd(nums[i], mx)
	}

	sort.Ints(prefixGcd)

	var ans int64
	left, right := 0, n-1

	for left < right {
		ans += int64(gcd(prefixGcd[left], prefixGcd[right]))
		left++
		right--
	}

	return ans
}

func main() {
	fmt.Println(gcdSum([]int{2, 6, 4}))    // 2
	fmt.Println(gcdSum([]int{3, 6, 2, 8})) // 5
}
