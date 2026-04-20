package main

import "fmt"

func maxDistance(colors []int) int {
	n := len(colors)

	maxDist := 0

	// Compare with first element
	for j := n - 1; j >= 0; j-- {
		if colors[j] != colors[0] {
			if j > maxDist {
				maxDist = j
			}
			break
		}
	}

	// Compare with last element
	for i := 0; i < n; i++ {
		if colors[i] != colors[n-1] {
			if (n - 1 - i) > maxDist {
				maxDist = n - 1 - i
			}
			break
		}
	}

	return maxDist
}

func main() {
	fmt.Println(maxDistance([]int{1, 1, 1, 6, 1, 1, 1})) // 3
	fmt.Println(maxDistance([]int{1, 8, 3, 8, 3}))       // 4
	fmt.Println(maxDistance([]int{0, 1}))                // 1
}
