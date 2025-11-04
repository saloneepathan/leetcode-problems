package main

import (
	"fmt"
	"sort"
)

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, 0, n-k+1)

	// helper function to calculate x-sum for current window
	calcXSum := func(freq map[int]int) int {
		type pair struct {
			val, count int
		}
		arr := make([]pair, 0, len(freq))
		for v, c := range freq {
			arr = append(arr, pair{v, c})
		}
		// sort by count desc, then value desc
		sort.Slice(arr, func(i, j int) bool {
			if arr[i].count == arr[j].count {
				return arr[i].val > arr[j].val
			}
			return arr[i].count > arr[j].count
		})

		sum := 0
		take := 0
		for _, p := range arr {
			if take == x {
				break
			}
			sum += p.val * p.count
			take++
		}
		return sum
	}

	freq := make(map[int]int)
	// initialize first window
	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}
	ans = append(ans, calcXSum(freq))

	// slide window
	for i := k; i < n; i++ {
		left := nums[i-k]
		freq[left]--
		if freq[left] == 0 {
			delete(freq, left)
		}
		right := nums[i]
		freq[right]++
		ans = append(ans, calcXSum(freq))
	}

	return ans
}

func main() {
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2)) // [6,10,12]
	fmt.Println(findXSum([]int{3, 8, 7, 8, 7, 5}, 2, 2))       // [11,15,15,15,12]
}
