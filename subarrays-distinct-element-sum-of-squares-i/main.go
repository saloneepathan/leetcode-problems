
package main

import "fmt"

func sumCounts(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		distinct := make(map[int]bool)

		for j := i; j < n; j++ {
			distinct[nums[j]] = true
			count := len(distinct)
			ans += count * count
		}
	}

	return ans
}

func main() {
	nums1 := []int{1, 2, 1}
	fmt.Println(sumCounts(nums1)) // Output: 15

	nums2 := []int{1, 1}
	fmt.Println(sumCounts(nums2)) // Output: 3
}
