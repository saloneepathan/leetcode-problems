package main

import "fmt"

func maximumLength(nums []int) int {
	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}

	ans := 1

	// Handle 1 separately since 1^2 = 1.
	if c, ok := cnt[1]; ok {
		if c%2 == 0 {
			ans = max(ans, c-1)
		} else {
			ans = max(ans, c)
		}
	}

	for x := range cnt {
		if x == 1 {
			continue
		}

		cur := x
		length := 0

		for {
			c, ok := cnt[cur]
			if !ok {
				break
			}

			if c >= 2 {
				length += 2
			} else {
				length++
				break
			}

			// Prevent overflow.
			if cur > 1000000000/cur {
				break
			}
			cur *= cur
		}

		// If no single peak exists, remove one element.
		if length > 0 && length%2 == 0 {
			length--
		}

		ans = max(ans, length)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumLength([]int{5, 4, 1, 2, 2}))  // 3
	fmt.Println(maximumLength([]int{1, 3, 2, 4}))     // 1
	fmt.Println(maximumLength([]int{2, 2, 4, 4, 16})) // 5
	fmt.Println(maximumLength([]int{1, 1, 1, 1}))     // 3
}
