package main

import (
	"fmt"
	"sort"
)

func minSwaps(nums []int) int {
	n := len(nums)

	digitSum := func(num int) int {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}

	type Pair struct {
		num int
		sum int
	}

	arr := make([]Pair, n)
	for i, num := range nums {
		arr[i] = Pair{
			num: num,
			sum: digitSum(num),
		}
	}

	// Sort by digit sum, then by number.
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].sum != arr[j].sum {
			return arr[i].sum < arr[j].sum
		}
		return arr[i].num < arr[j].num
	})

	// Map each number to its target index.
	target := make(map[int]int, n)
	for i, p := range arr {
		target[p.num] = i
	}

	// Count permutation cycles.
	visited := make([]bool, n)
	swaps := 0

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		current := i
		cycleSize := 0

		for !visited[current] {
			visited[current] = true
			current = target[nums[current]]
			cycleSize++
		}

		if cycleSize > 1 {
			swaps += cycleSize - 1
		}
	}

	return swaps
}

func main() {
	nums1 := []int{37, 100}
	fmt.Println(minSwaps(nums1)) // Output: 1

	nums2 := []int{22, 14, 33, 7}
	fmt.Println(minSwaps(nums2)) // Output: 0

	nums3 := []int{18, 43, 34, 16}
	fmt.Println(minSwaps(nums3)) // Output: 2
}