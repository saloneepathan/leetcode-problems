package main

import (
	"fmt"
)

func longestBalanced(nums []int) int {
	n := len(nums)
	ans := 0

	for l := 0; l < n; l++ {
		evenFreq := make(map[int]int)
		oddFreq := make(map[int]int)
		distinctEven, distinctOdd := 0, 0

		for r := l; r < n; r++ {
			num := nums[r]

			if num%2 == 0 {
				if evenFreq[num] == 0 {
					distinctEven++
				}
				evenFreq[num]++
			} else {
				if oddFreq[num] == 0 {
					distinctOdd++
				}
				oddFreq[num]++
			}

			if distinctEven == distinctOdd {
				if r-l+1 > ans {
					ans = r - l + 1
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(longestBalanced([]int{2, 5, 4, 3}))    // Expected: 4
	fmt.Println(longestBalanced([]int{3, 2, 2, 5, 4})) // Expected: 5
	fmt.Println(longestBalanced([]int{1, 2, 3, 2}))    // Expected: 3
	fmt.Println(longestBalanced([]int{22, 36, 22}))    // Expected: 0
}
